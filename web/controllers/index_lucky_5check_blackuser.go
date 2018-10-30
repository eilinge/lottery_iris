package controllers

import (
	"time"
	"imooc.com/lottery/models"
)

// 验证当前用户是否存在黑名单限制
func (c *IndexController) checkBlackUser(uid int) (bool, *models.LtUser) {
	info := c.ServiceUser.Get(uid)
	if info != nil && info.Blacktime > int(time.Now().Unix()) {
		// IP黑名单存在，而且没有过期
		return false, info
	}
	return true, info
}
