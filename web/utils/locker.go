package utils

import (
	"fmt"
	"lottery/datasource"
)

/* 防止同一用户的重入
不安全: 并发验证, 并发数据更新
分布式锁: 保证原子性, 避免死锁
用户请求结束, 必须即使释放锁
*/

func getLuckyLockKey(uid int) string {
	return fmt.Sprintf("lucky_lock_%d", uid)
}

// LockLucky lock
func LockLucky(uid int) bool {
	key := getLuckyLockKey(uid)
	cacheObj := datasource.InstanceCache()
	// EX 设置过期时间, NX 是否存在
	rs, _ := cacheObj.Do("SET", key, 1, "EX", 3, "NX")
	if rs == "OK" {
		return true
	}
	return false
}

// UnLockLucky unlock
func UnLockLucky(uid int) bool {
	key := getLuckyLockKey(uid)
	cacheObj := datasource.InstanceCache()
	rs, _ := cacheObj.Do("DEL", key)
	if rs == "OK" {
		return true
	}
	return false
}