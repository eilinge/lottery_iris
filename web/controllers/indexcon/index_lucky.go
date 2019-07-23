package indexcon

import (
	"log"

	"lottery/comm"
	"lottery/conf"
	"lottery/models"
	"lottery/web/utils"
)

// GetLucky 抽奖入口
func (c *IndexController) GetLucky() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = ""

	// 1. 验证登录用户
	loginuser := comm.GetLoginUser(c.Ctx.Request())
	if loginuser == nil || loginuser.Uid < 1 {
		rs["code"] = 101
		rs["msg"] = "请先登录, 再来抽奖"
		return rs
	}

	// 2. 用户抽奖分布式锁定
	ok := utils.LockLucky(loginuser.Uid)
	if ok {
		defer utils.UnLockLucky(loginuser.Uid)
	} else {
		rs["code"] = 102
		rs["msg"] = "正在抽奖, 请稍候重试"
	}

	// 3. 验证用户进入参与次数
	/*
		强制限制同一个用户/IP, 不能超过今天参与次数
		每日参与次数的数据更新, 原子性递增
		IP的今日参与次数需要用redis缓存来实现
	*/
	userDayNum := utils.IncrUserLuckNum(loginuser.Uid)
	if userDayNum > conf.UserPrizeMax {
		rs["code"] = 103
		rs["msg"] = "今日的抽奖次数已用完, 明天再来吧"
		return rs
	}
	ok = c.checkUserday(loginuser.Uid)
	if !ok{
		rs["code"] = 103
		rs["msg"] = "今日的抽奖次数已用完, 明天再来吧"
		return rs
	}
	// 4. 验证IP今日的参与次数
	ip := comm.ClientIP(c.Ctx.Request())
	ipDayNum := utils.IncrIPLuckNum(ip)
	if ipDayNum > conf.IpLimitMax {
		rs["code"] = 104
		rs["msg"] = "相同IP参与次数太多, 明天再来参与吧"
		return rs
	}
	// 5. 验证IP黑名单
	/* 黑名单不能获取实物奖
		安全机制, 避免同一个ip大量用户的刷奖
		公平机制, 中过大奖的用户在一段时间内把机会让出来
		验证方法, 数据存在并且在黑名单限制单位时间内
	*/
	limitBlack := false
	if ipDayNum > conf.IpPrizeMax {
		limitBlack = true
	}
	var blackipInfo *models.LtBlackip
	if !limitBlack {
		ok, _ := c.checkBlackip(ip)
		// 验证不通过
		if !ok {
			log.Println("黑名单的IP", ip, limitBlack)
			limitBlack = true
		}
	}
	// 6. 验证用户黑名单
	var userInfo *models.LtUser
	if !limitBlack {
		ok, _ := c.checkBlackUser(loginuser.Uid)
		if !ok {
			log.Println("黑名单的用户uid", loginuser.Uid, limitBlack)
			limitBlack = true
		}
	}
	// 7. 获取抽奖编码
	prizeCode := comm.Random(10000)
	// 8. 匹配奖品是否中奖
	prizeGift := c.prize(prizeCode, limitBlack)
	if prizeGift == nil ||
		prizeGift.PrizeNum < 0 || (prizeGift.PrizeNum > 0 && prizeGift.LeftNum < 0) {
			rs["code"] = 205
			rs["msg"] = "很遗憾, 没有中奖, 请下次再试"
			return rs
	}
	// 9. 有限制奖品发放
	if prizeGift.PrizeNum > 0 {
		ok = utils.PrizeGift(prizeGift.Id, prizeGift.LeftNum)
		if !ok {
			rs["code"] = 207
			rs["msg"] = "很遗憾, 没有中奖, 请下次再试"
			return rs
		}
	}
	// 10. 不同编码的优惠券的发放
	if prizeGift.Gtype == conf.GtypeCodeDiff {
		code := utils.PrizeCodeDiff(prizeGift.Id, c.ServiceCode)
		if code == "" {
			rs["code"] = 208
			rs["msg"] = "很遗憾, 没有中奖, 请下次再试"
			return rs
		}
		prizeGift.Gdata = code
	}
	// 11. 记录中奖纪录
	result := models.LtResult{
		GiftId: prizeGift.Id,
		GiftName: prizeGift.Title,
		GiftType: prizeGift.Gtype,
		Uid: loginuser.Uid,
		Username: loginuser.Username,
		PrizeCode: prizeCode,
		GiftData: prizeGift.Gdata,
		SysCreated: comm.NowUnix(),
		SysIp: ip,
		SysStatus: 0,
	}
	err := c.ServiceResult.Create(&result)
	if err != nil {
		log.Println("index_lucky.GetLucky ServiceResult.Create ", result, ", err=: ", err)
		rs["code"] = 209
		rs["msg"] = "很遗憾, 没有中奖, 请下次再试"
		return rs
	}
	// 12. 返回中奖结果
	if prizeGift.Gtype == conf.GtypeGiftLarge {
		// 如果获得实物大奖, 需要将用户, IP设置成黑名单一段时间
		c.prizeLarge(ip, loginuser, userInfo, blackipInfo)
	}
	rs["gift"] = prizeGift
	return rs
}