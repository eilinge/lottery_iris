/**
 * 首页根目录的Controller
 * http://localhost:8080/
 */
package controllers

import (
	"imooc.com/lottery/comm"
	"imooc.com/lottery/web/utils"
	"fmt"
	"imooc.com/lottery/conf"
	"imooc.com/lottery/models"
	"log"
)

// http://localhost:8080/lucky
func (c *IndexController) GetLucky() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = ""
	// 1 TODO: 验证登录 (cookie验证)
	loginuser := comm.GetLoginUser(c.Ctx.Request())
	//if loginuser == nil || loginuser.Uid < 1 {
	//	rs["code"] = 101
	//	rs["msg"] = "请先登录，再来抽奖"
	//	return rs
	//}
	if loginuser == nil || loginuser.Uid < 1 {
		loginuser = &models.ObjLoginuser{
			Uid:      comm.Random(10000),
			Username: "admin",
			Now:      0,
			Ip:       ":1",
			Sign:     "abc",
		}
	}

	// 开始抽奖的业务逻辑 =======================
	// 2 用户抽奖分布式锁定 (锁，使用本地或者redis)
	ok := utils.LockLucky(loginuser.Uid)
	if ok { // 加锁成功才可以继续，而且要注意最后解锁
		defer utils.UnlockLucky(loginuser.Uid)
	} else {
		rs["code"] = 102
		rs["msg"] = "正在抽奖，请稍等重试"
		return rs
	}

	// 3 验证用户今日参与次数 (redis或者redis)
	// 需要从数据库中获取用户当天的数据
	ok = c.checkUserday(loginuser.Uid)
	if !ok {
		rs["code"] = 103
		rs["msg"] = "今天的抽奖次数已用完，明天再来吧"
		return rs
	}

	// 4 验证同一个IP当天的限制次数
	ip := comm.ClientIP(c.Ctx.Request())
	ipDaynum := utils.IncrIpLucyNum(ip)
	if ipDaynum > conf.IpLimitMax {
		rs["code"] = 104
		rs["msg"] = "相同IP参与次数太多，明天再来参与吧"
		return rs
	}

	limitBlack := false	// 黑名单信息，不能抽实物奖
	// 5 验证IP今日参与次数
	if ipDaynum > conf.IpPrizeMax {
		limitBlack = true
	}

	// 6 验证IP黑名单
	var blackipInfo *models.LtBlackip
	if !limitBlack {
		ok, blackipInfo = c.checkBlackip(ip)
		if !ok {
			fmt.Println("黑名单中的IP", ip, limitBlack)
			limitBlack = true
		}
	}

	var userInfo *models.LtUser
	// 7 验证用户黑名单
	if !limitBlack {
		ok, userInfo = c.checkBlackUser(loginuser.Uid)
		if !ok {
			fmt.Println("黑名单中的用户", loginuser.Uid, limitBlack)
			limitBlack = true
		}
	}

	// 8 获得抽奖编码
	prizeCode := comm.Random(10000)

	// 9 匹配奖品是否中奖
	prizeGift := c.prize(prizeCode, limitBlack)
	if prizeGift == nil ||
		prizeGift.PrizeNum < 0 ||
		(prizeGift.PrizeNum > 0 && prizeGift.LeftNum <= 0) {
		rs["code"] = 205
		rs["msg"] = "很遗憾，没有中奖，请下次再试"
		return rs
	}

	log.Println("index_lucky.GetLucky prizeCode=", prizeCode, ", gift=", prizeGift)
	// 10.1 有限制奖品发放
	if prizeGift.PrizeNum > 0 {
		// 10.2 有限制奖品发放（奖品池中剩余数量）
		// 奖品池的数据处理和更新
		ok = utils.PrizeGift(prizeGift.Id, prizeGift.LeftNum)
		if !ok {
			rs["code"] = 207
			rs["msg"] = "很遗憾，没有中奖，请下次再试"
			return rs
		}

		// 10.3 不同编码的虚拟券发放
		if prizeGift.Gtype == conf.GtypeCodeDiff {
			prizeGift.Gdata = utils.PrizeCodeDiff(prizeGift.Id, c.ServiceCode)
			if prizeGift.Gdata == "" {
				// 没有得到一个正常的优惠券编码
				rs["code"] = 208
				rs["msg"] = "很遗憾，没有中奖，请下次再试"
				return rs
			}
		}
	}

	// 11 记录中奖纪录
	result := models.LtResult{
		GiftId:     prizeGift.Id,
		GiftName:   prizeGift.Title,
		GiftType:   prizeGift.Gtype,
		Uid:        loginuser.Uid,
		Username:   loginuser.Username,
		PrizeCode:  prizeCode,
		GiftData:   prizeGift.Gdata,
		SysCreated: comm.NowUnix(),
		SysIp:      ip,
		SysStatus:  0,
	}
	err := c.ServiceResult.Create(&result)
	if err != nil {
		log.Println("index_lucky.GetLucky ServiceResult.Create ", result, ", error=", err)
		rs["code"] = 209
		rs["msg"] = "很遗憾，没有中奖，请下次再试"
		return rs
	}
	if prizeGift.Gtype == conf.GtypeGiftLarge {
		// 如果获得了实物大奖，需要将用户、IP设置成黑名单一段时间
		c.prizeLarge(ip, loginuser, userInfo, blackipInfo)
	}

	// 12 返回抽奖结果
	rs["gift"] = prizeGift
	return rs
}
