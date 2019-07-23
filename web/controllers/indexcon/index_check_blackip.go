package indexcon

import (
	"time"

	"lottery/models"
)

func (c *IndexController) checkBlackip(ip string) (bool, *models.LtBlackip) {
	info := c.ServiceBlackip.GetByIp(ip)
	if info == nil || info.Ip == "" {
		return true, nil
	}
	if info.Blacktime > int(time.Now().Unix()) {
		return false, info
	}

	return true, info
}