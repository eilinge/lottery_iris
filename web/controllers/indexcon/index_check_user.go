package indexcon

import (
	"time"

	"lottery/models"
)
func (c *IndexController) checkBlackUser(uid int) (bool, *models.LtUser) {
	info := c.ServiceUser.Get(uid)
	if info != nil && info.Blacktime > int(time.Now().Unix()) {
		return false, info
	}

	return true, info
}