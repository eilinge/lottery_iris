package utils

import (
	"log"

	"lottery/comm"
	"lottery/services"
)
// PrizeGift 对奖品剩余数递减
func PrizeGift(id, leftNum int) bool{
	giftService := services.NewGiftService()
	rows, err := giftService.DecrLeftNum(id, 1)
	if rows < 1 || err != nil {
		log.Println("prizedata giftService.DescleftNum err: ", err, ", row=", rows)
		return false
	}
	return true
}

// PrizeCodeDiff 不同奖品的优惠券
func PrizeCodeDiff(id int, codeService services.CodeService) string {
	lockUid := 0 - id - 100000000
	LockLucky(lockUid)
	defer UnLockLucky(lockUid)

	codeId := 0
	codeInfo := codeService.NextUsingCode(id, codeId)
	if codeInfo != nil && codeInfo.Id > 0{
		codeInfo.SysStatus = 2
		codeInfo.SysUpdated = comm.NowUnix()
		codeService.Update(codeInfo, nil)
	} else {
		log.Println("pirzedata.PrizeCodeDiff num codeInfo, gift_id = ", codeId)
		return ""
	}
	return ""
}
