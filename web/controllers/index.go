/**
 * 首页根目录的Controller
 * http://localhost:8080/
 */
package controllers

import (
	"github.com/kataras/iris"
			"fmt"

	"imooc.com/lottery/services"
	"imooc.com/lottery/models"
	"imooc.com/lottery/comm"
	)

type IndexController struct {
	Ctx iris.Context
	ServiceUser services.UserService
	ServiceGift services.GiftService
	ServiceCode services.CodeService
	ServiceResult services.ResultService
	ServiceUserday services.UserdayService
	ServiceBlackip services.BlackipService
}

// http://localhost:8080/
func (c *IndexController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return "welcome to Go抽奖系统，<a href='/public/index.html'>开始抽奖</a>"
}

// http://localhost:8080/gifts
func (c *IndexController) GetGifts() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = ""
	datalist := c.ServiceGift.GetAll()
	list := make([]models.LtGift, 0)
	for _, data := range datalist {
		// 正常状态的才需要放进来
		if data.SysStatus == 0 {
			list = append(list, data)
		}
	}
	rs["gifts"] = list
	return rs
}

// http://localhost:8080/newprize
func (c *IndexController) GetNewprize() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = ""
	gifts := c.ServiceGift.GetAll()
	giftIds := []int{}
	for _, data := range gifts {
		// 虚拟券或者实物奖才需要放到外部榜单中展示
		if data.Gtype > 1 {
			giftIds = append(giftIds, data.Id)
		}
	}
	list := c.ServiceResult.GetNewPrize(50, giftIds)
	rs["prize_list"] = list
	return rs
}

// 登录 GET /login
func (c *IndexController) GetLogin() {
	// 每次随机生成一个登录用户信息
	uid := comm.Random(100000)
	loginuser := models.ObjLoginuser{
		Uid:      uid,
		Username: fmt.Sprintf("admin-%d", uid),
		Now:      comm.NowUnix(),
		Ip:       comm.ClientIP(c.Ctx.Request()),
	}
	comm.SetLoginuser(c.Ctx.ResponseWriter(), &loginuser)
	comm.Redirect(c.Ctx.ResponseWriter(), "/public/index.html?from=login")
}

// 退出 GET /logout
func (c *IndexController) GetLogout() {
	comm.SetLoginuser(c.Ctx.ResponseWriter(), nil)
	comm.Redirect(c.Ctx.ResponseWriter(), "/public/index.html?from=logout")
}
