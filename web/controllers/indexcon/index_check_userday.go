package indexcon

import (
	"log"
	"time"
	"strconv"
	"fmt"

	"lottery/models"
	"lottery/conf"
)

func (c *IndexController) checkUserday(uid int) bool {
	userdayInfo := c.ServiceUserday.GetUserToday(uid)
	if userdayInfo != nil && userdayInfo.Uid == uid {
		// 今天存在抽奖记录
		if userdayInfo.Num >= conf.UserPrizeMax {
			return false
		}
		userdayInfo.Num++
		err := c.ServiceUserday.Update(userdayInfo, nil)
		if err != nil {
			log.Println("failed to ServiceUserday.Update err: ", err)
		}
	} else {
		// 创建今天的用户参与记录
		y, m, d := time.Now().Date()
		strDay := fmt.Sprintf("%d%02d%02d", y, m, d)
		day, _ := strconv.Atoi(strDay)
		userdayInfo = &models.LtUserday{
			Uid: uid,
			Day: day,
			Num: 1,
			SysCreated: int(time.Now().Unix()),
		}
		err := c.ServiceUserday.Create(userdayInfo)
		if err != nil {
			log.Println("failed to ServiceUserday.Update err: ", err)
		}
	}
	return true
}