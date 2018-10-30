package controllers

import (
	"imooc.com/lottery/models"
	"imooc.com/lottery/comm"
)

// 中了实物大奖，需要把用户、IP暂时放入到黑名单一段时间，人品暂时用光了
func (c *IndexController) prizeLarge(ip string,
	loginuser *models.ObjLoginuser,
	userInfo *models.LtUser,
	blackipInfo *models.LtBlackip) {
	nowTime := comm.NowUnix()
	blackTime := 30 * 86400
	// 更新用户的黑名单信息
	if userInfo == nil || userInfo.Id <= 0 {
		userInfo = &models.LtUser{
			Id:         loginuser.Uid,
			Username:   loginuser.Username,
			Blacktime:  nowTime + blackTime,
			SysCreated: nowTime,
			SysIp:      ip,
		}
		c.ServiceUser.Create(userInfo)
	} else {
		userInfo = &models.LtUser{
			Id:loginuser.Uid,
			Blacktime:nowTime + blackTime,
			SysUpdated:nowTime,
		}
		c.ServiceUser.Update(userInfo, nil)
	}
	// 更新IP的黑名单信息
	if blackipInfo == nil || blackipInfo.Id <= 0 {
		blackipInfo = &models.LtBlackip{
			Ip:         ip,
			Blacktime:  nowTime + blackTime,
			SysCreated: nowTime,
		}
		c.ServiceBlackip.Create(blackipInfo)
	} else {
		blackipInfo.Blacktime = nowTime + blackTime
		blackipInfo.SysUpdated = nowTime
		c.ServiceBlackip.Update(blackipInfo, nil)
	}
}
