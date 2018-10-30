package controllers

import (
	"time"
	"strconv"
	"log"

	"imooc.com/lottery/models"
	"imooc.com/lottery/conf"
	"imooc.com/lottery/web/utils"
	"fmt"
)

// 验证用户今天的抽奖次数是否超过每天最大抽奖次数
func (c *IndexController) checkUserday(uid int, num int64) bool {
	// index_lucky.go 中 utils.IncrUserLuckyNum(loginuser.Uid) 从缓存中验证
	// 缓存验证后，还可以继续抽奖，才需要在这里继续处理
	userdayInfo := c.ServiceUserday.GetUserToday(uid)
	if userdayInfo != nil && userdayInfo.Uid == uid {
		// 今天存在抽奖记录
		if userdayInfo.Num >= conf.UserPrizeMax {
			// 缓存中的今日参与次数太小，需要重新更新到缓存
			if int(num) < userdayInfo.Num {
				utils.InitUserLuckyNum(uid, int64(userdayInfo.Num))
			}
			return false
		} else {
			// 更新今天的抽奖次数
			userdayInfo.Num++
			// 缓存中的今日参与次数太小，需要重新更新到缓存
			if int(num) < userdayInfo.Num {
				utils.InitUserLuckyNum(uid, int64(userdayInfo.Num))
			}
			// 更新到数据库
			err103 := c.ServiceUserday.Update(userdayInfo, nil)
			if err103 != nil {
				log.Println("index_lucky_check_userday ServiceUserday.Update err103=", err103)
			}
		}
	} else {
		// 创建今天的用户参与记录
		y, m, d := time.Now().Date()
		strDay := fmt.Sprintf("%d%02d%02d", y, m, d)
		day, _ := strconv.Atoi(strDay)
		userdayInfo = &models.LtUserday{
			Uid:        uid,
			Day:        day,
			Num:        1,
			SysCreated: int(time.Now().Unix()),
		}
		err103 := c.ServiceUserday.Create(userdayInfo)
		if err103 != nil {
			log.Println("index_lucky_check_userday ServiceUserday.Create err103=", err103)
		}
		// 用户今天的第一次抽奖
		utils.InitUserLuckyNum(uid, 1)
	}
	return true
}
