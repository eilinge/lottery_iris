package controllers

import (
	"imooc.com/lottery/models"
	"imooc.com/lottery/conf"
	)

func (c *IndexController) prize(prizeCode int, limitBlack bool) *models.ObjGiftPrize {
	var prizeGift *models.ObjGiftPrize
	giftList := c.ServiceGift.GetAllUse(true)

	for _, gift := range giftList {
		if gift.PrizeCodeA <= prizeCode && gift.PrizeCodeB >= prizeCode {
			// 中奖编码满足条件，说明可以中奖
			if !limitBlack || gift.Gtype < conf.GtypeGiftSmall {
				// 不在黑名单限制，或者，不是实物奖
				prizeGift = &gift
				break
			}
		}
	}
	return prizeGift
}
