/**
 * 首页根目录的Controller
 * http://localhost:8080/
 */
package controllers

import (
	"github.com/kataras/iris"
	"fmt"
	"lottery/models"
	"lottery/services"
	"lottery/comm"
)

// IndexController 其他用户访问界面
type IndexController struct {
	Ctx            iris.Context // 解析前端传来的数据
	ServiceUser    services.UserService
	ServiceGift    services.GiftService
	ServiceCode    services.CodeService
	ServiceResult  services.ResultService
	ServiceUserday services.UserdayService
	ServiceBlackip services.BlackipService
}

// Get http://localhost:8080/
func (c *IndexController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return "Welcome to Go抽奖系统, <a href='/public/index.html'>开始抽奖</a>"
}

// GetGifts http://localhost:8080/gifts
func (c *IndexController) GetGifts() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = ""
	datalist := c.ServiceGift.GetAll(false)

	list := make([]models.LtGift, 0)

	for _, data := range datalist {
		if data.SysStatus == 0 {
			list = append(list, data)
		}
	}
	rs["gifts"] = list
	return rs
}

// GetNewprize http://localhost:8080/newprize
func (c *IndexController) GetNewprize() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = ""
	// TODO:
	return rs
}

// GetLogin http://localhost:8080/login
func (c *IndexController) GetLogin() {
	uid := comm.Random(100000)
	loginuser := models.ObjLoginuser{
		Uid: uid,
		Username: fmt.Sprintf("admin-%d", uid),
		Now: comm.NowUnix(),
		Ip: comm.ClientIP(c.Ctx.Request()),
	}

	comm.SetLoginuser(c.Ctx.ResponseWriter(), &loginuser)
	comm.Redirect(c.Ctx.ResponseWriter(), 
		"/public/index.html?from=login")
}

// GetLogout http://localhost:8080/logout
func (c *IndexController) GetLogout() {
	comm.SetLoginuser(c.Ctx.ResponseWriter(), nil)
	comm.Redirect(c.Ctx.ResponseWriter(),
		"/public/index.html?from=logout")
}
// GetClearcache ...
// 集群多服务器的时候，才用得上这个接口
// 性能优化的时候才考虑，加上本机的SQL缓存
// http://localhost:8080/clearcache
/*
func (c *IndexController) GetClearcache() mvc.Result {
	// *xorm.Engine.ClearCache()
	err := datasource.InstanceMaster().ClearCache(&models.StarInfo{})
	if err != nil {
		log.Fatal(err)
	}
	// set the model and render the view template.
	return mvc.Response{
		Text: "xorm缓存清除成功",
	}
}
*/
