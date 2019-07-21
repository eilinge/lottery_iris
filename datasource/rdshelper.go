package datasource

import (
	"fmt"
	"log"
	"lottery/conf"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

var rdsLock sync.Mutex
var cacheInstance *RedisConn

type RedisConn struct {
	Pool      *redis.Pool
	showDebug bool
}

func InstanceCache() *RedisConn {
	if cacheInstance != nil {
		return cacheInstance
	}
	rdsLock.Lock()
	defer rdsLock.Unlock()
	if cacheInstance != nil {
		return cacheInstance
	}

	return NewCache()
}

func (rds *RedisConn) Do(command string, args ...interface{}) (reply interface{}, err error) {
	conn := rds.Pool.Get()
	defer conn.Close() // 该连接执行完之后, 放入连接池中

	t1 := time.Now().UnixNano()
	reply, err = conn.Do(command, args...)
	if err != nil {
		e := conn.Err()
		if e != nil {
			log.Println("rdshelper.Do err: ", err, e)
		}
	}

	t2 := time.Now().UnixNano()
	if rds.showDebug {
		fmt.Printf("[redis] [info] [%dus] cmd=%s, err=%s, args=%v, reply=%s\n",
			(t2-t1)/1000, command, err, args, reply)
	}
	return reply, err
}

func (rds *RedisConn) ShowDebug(b bool) {
	rds.showDebug = b
}

func NewCache() *RedisConn {
	pool := redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", conf.RdsCache.Host, conf.RdsCache.Port))
			if err != nil {
				log.Fatal("rdshelper.NewCache Dial error: ", err)
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:         10000, // 最大连接数
		MaxActive:       10000, // 最大活跃数
		IdleTimeout:     0,     // 超时时间
		Wait:            false,
		MaxConnLifetime: 0, // 活跃时间
	}
	instance := &RedisConn{
		Pool: &pool,
	}
	cacheInstance = instance
	instance.ShowDebug(true)
	return instance
}
