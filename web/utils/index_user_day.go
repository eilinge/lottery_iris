package utils

import (
	"fmt"
	"math"
	"log"
	"time"

	"lottery/datasource"
	"lottery/comm"
)

const userFrameSize = 2

func init() {
	resetGroupUserList()
}
/*
用户今日抽奖次数, hash中的计数器, utils/user_day

incrUserLuckyNum 原子性递增用户今日的抽奖次数
initUserLuckyNum 以数据库数据为准, 从数据库初始化缓存数据(单个用户抽奖次数需要很准确)

resetGroupUserList 每天凌晨计数器归零

userFrameSize 优化, 将hash结构散列为多段数据, 让每个hash小点, 以提高redis的执行效率
*/
func resetGroupUserList() {
	log.Println("user_day_lucky.resetGroupUserList start")
	cacheObj := datasource.InstanceCache()
	for i := 0; i < userFrameSize; i++ {
		key := fmt.Sprintf("day_users_%d", i)
		cacheObj.Do("DEL", key)
	}
	log.Println("user_day_lucky.resetGroupUserList stop")
	// user当天的统计数, 零点的时候归零, 设置定时器
	duration := comm.NextDayDuration()
	time.AfterFunc(duration, resetGroupUserList)
}

// IncrUserLuckNum 放入缓存中, 进行累加
func IncrUserLuckNum(uid int) int64 {
	i := uid % userFrameSize
	key := fmt.Sprintf("day_users_%d", i)
	cacheObj := datasource.InstanceCache()
	rs, err := cacheObj.Do("HINCRBY", key, uid, 1)
	if err != nil {
		log.Println("uid_day_lucky redis HINCRBY error=", err)
		return math.MaxInt32
	}
	return rs.(int64)
}

// InitUserLuckNum ...
func InitUserLuckNum(uid int, num int64) {
	if num <= 1 {
		return
	}
	i := uid % userFrameSize
	key := fmt.Sprintf("day_users_%d", i)
	cacheObj := datasource.InstanceCache()
	_, err := cacheObj.Do("HSET", key, uid, num)
	if err != nil {
		log.Println("user_day.InitUserLuckNum redis HSET key=", key, ",error=", err)
		return
	}
}