package utils

import (
	"fmt"
	"math"
	"log"
	"time"

	"lottery/datasource"
	"lottery/comm"
)

const ipFrameSize = 2

func init() {
	resetGroupIPList()
}

func resetGroupIPList() {
	log.Println("ip_day_lucky.resetGroupIPList start")
	cacheObj := datasource.InstanceCache()
	for i := 0; i < ipFrameSize; i++ {
		key := fmt.Sprintf("day_ips_%d", i)
		cacheObj.Do("DEL", key)
	}
	log.Println("ip_day_lucky.resetGroupIPList stop")
	// ip当天的统计数, 零点的时候归零, 设置定时器
	duration := comm.NextDayDuration()
	time.AfterFunc(duration, resetGroupIPList)
}

// IncrIPLuckNum 放入缓存中, 进行累加
func IncrIPLuckNum(strIP string) int64 {
	ip := comm.IP4toInt(strIP)
	i := ip % ipFrameSize
	key := fmt.Sprintf("day_ips_%d", i)
	cacheObj := datasource.InstanceCache()
	rs, err := cacheObj.Do("HINCRBY", key, ip, 1)
	if err != nil {
		log.Println("ip_day_lucky redis HINCRBY error=", err)
		return math.MaxInt32
	}
	return rs.(int64)
}