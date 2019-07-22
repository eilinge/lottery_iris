/**
 * 首页根目录的Controller
 * http://localhost:8080/
 */
package controllers

import (
	"fmt"
	"log"
	"lottery/comm"
	"lottery/models"
	"lottery/services"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// AdminBlackipController 其他用户访问界面
type AdminBlackipController struct {
	Ctx            iris.Context // 解析前端传来的数据
	ServiceBlackip services.BlackipService
}

// Get http://localhost:8080/
func (c *AdminBlackipController) Get() mvc.Result {
	page := c.Ctx.URLParamIntDefault("page", 1)

	size := 100
	pagePrev := ""
	pageNext := ""

	dataList := c.ServiceBlackip.GetAll(page, size)

	total := (page - 1) + len(dataList)
	// 数据总数
	if len(dataList) >= size {
		total = int(c.ServiceBlackip.CountAll())
		pageNext = fmt.Sprintf("%d", page+1)
	}
	if page > 1 {
		pagePrev = fmt.Sprintf("%d", page-1)
	}

	log.Println("dataList： ", dataList)
	return mvc.View{
		Name: "admin/blackip.html",
		Data: iris.Map{
			"Title":    "管理后台",
			"Channel":  "Blackip",
			"Datalist": dataList,
			"Total":    total,
			"Now":      comm.NowUnix(),
			"PagePrev": pagePrev,
			"PageNext": pageNext,
		},
		Layout: "admin/layout.html",
	}
}

// GetBlack http://localhost:8080/admin/Blackip/black?id=1&time=0
func (c *AdminBlackipController) GetBlack() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	t := c.Ctx.URLParamIntDefault("time", 0)

	if err == nil {
		if t > 0 {
			t = t*86400 + comm.NowUnix()
		}
		c.ServiceBlackip.Update(&models.LtBlackip{Id: id,
			Blacktime: t, SysUpdated: comm.NowUnix()},
			[]string{"blacktime"})
	}
	return mvc.Response{
		Path: "/admin/blackip",
	}
}
