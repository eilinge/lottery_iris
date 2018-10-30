package utils

import (
		"fmt"
	"log"
		"imooc.com/lottery/comm"
		"imooc.com/lottery/datasource"
	"imooc.com/lottery/models"
	"imooc.com/lottery/services"
)

// 发奖，指定的奖品是否还可以发出来奖品
func PrizeGift(id, leftNum int) bool {
	// 更新数据库，减少奖品的库存
	giftService := services.NewGiftService()
	rows, err := giftService.DecrLeftNum(id, 1)
	if rows < 1 || err != nil {
		log.Println("prizedata.PrizeGift giftService.DecrLeftNum error=", err, ", rows=", rows)
		// 数据更新失败，不能发奖
		return false
	}
	return true
}

// 优惠券类的发放
func PrizeCodeDiff(id int, codeService services.CodeService) string {
	//return prizeServCodeDiff(id, codeService)
	return prizeLocalCodeDiff(id, codeService)
}

// 优惠券发放，只使用数据库的方式发放，手动加锁
func prizeLocalCodeDiff(id int, codeService services.CodeService) string {
	//muPrizeCode.Lock()
	//defer muPrizeCode.Unlock()
	// 复用用户抽奖的锁
	lockUid := 0-id-100000000
	LockLucky(lockUid)
	defer UnlockLucky(lockUid)

	codeId := 0
	//// 从数据库中找到最小的未发放的券
	//codeId, ok := minPrizeCodeId[id]
	//if !ok {
	//	codeId = 0
	//}
	codeInfo := codeService.NextUsingCode(id, codeId)
	// 能否成功设置这个券发放状态
	if codeInfo != nil && codeInfo.Id > 0 {
		// 状态设置为已发放
		codeInfo.SysStatus = 2
		codeInfo.SysUpdated = comm.NowUnix()
		codeService.Update(codeInfo, nil)
	} else {
		log.Println("prizedata.prizeLocalCodeDiff null codeInfo, gift_id=", id)
		return ""
	}
	//minPrizeCodeId[id] = codeInfo.Id
	return codeInfo.Code
}

// 优惠券发放，使用redis的方式发放
func prizeServCodeDiff(id int, codeService services.CodeService) string {
	key := fmt.Sprintf("gift_code_%d", id)
	cacheObj := datasource.InstanceCache()
	rs, err := cacheObj.Do("SPOP", key)
	if err != nil {
		log.Println("prizedata.prizeServCodeDiff error=", err)
		return ""
	}
	code := comm.GetString(rs, "")
	if code == "" {
		log.Printf("prizedata.prizeServCodeDiff rs=%s", rs)
		return ""
	}
	// 更新数据库中的发放状态
	codeService.UpdateByCode(&models.LtCode{
		Code:       code,
		SysStatus:  2,
		SysUpdated: comm.NowUnix(),
	}, nil)
	return code
}
