package indexcon

import (
	"log"
	"time"
	"strconv"
	"fmt"

	"lottery/models"
	"lottery/conf"
	"lottery/web/utils"
)

func (c *IndexController) checkUserday(uid int, num int64) bool {
	userdayInfo := c.ServiceUserday.GetUserToday(uid)
	if userdayInfo != nil && userdayInfo.Uid == uid {
		// 今天存在抽奖记录
		if userdayInfo.Num >= conf.UserPrizeMax {
			if int(num) < userdayInfo.Num {
				utils.InitUserLuckNum(uid, int64(userdayInfo.Num))
			}
			return false
		}
		userdayInfo.Num++
		if int(num) < userdayInfo.Num {
				utils.InitUserLuckNum(uid, int64(userdayInfo.Num))
			}
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
		utils.InitUserLuckNum(uid, 1)
	}
	return true
}