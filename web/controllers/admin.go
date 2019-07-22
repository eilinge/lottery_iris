/**
 * 首页根目录的Controller
 * http://localhost:8080/
 */
package controllers

import (
	"github.com/kataras/iris"

	"lottery/models"
	"lottery/services"
	"github.com/kataras/iris/mvc"
)

// AdminController 其他用户访问界面
type AdminController struct {
	Ctx            iris.Context // 解析前端传来的数据
	ServiceUser    services.UserService
	ServiceGift    services.GiftService
	ServiceCode    services.CodeService
	ServiceResult  services.ResultService
	ServiceUserday services.UserdayService
	ServiceBlackip services.BlackipService
}

// Get http://localhost:8080/
func (c *AdminController) Get() mvc.Result {
	return mvc.View{
		Name: "admin/index.html",
		Data: iris.Map{
			"Title": "管理后台",
			"Channel": "",
		},
		Layout: "admin/layout.html",
	}
}

// GetGifts http://localhost:8080/gifts
func (c *AdminController) GetGifts() map[string]interface{} {
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
func (c *AdminController) GetNewprize() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = ""
	// TODO:
	return rs
}

