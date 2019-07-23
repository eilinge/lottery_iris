/**
 * 首页根目录的Controller
 * http://localhost:8080/
 */
package admincon

import (
	"fmt"
	"log"
	"lottery/models"
	"lottery/services"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// AdminResultController 其他用户访问界面
type AdminResultController struct {
	Ctx           iris.Context // 解析前端传来的数据
	ServiceGift   services.GiftService
	ServiceResult services.ResultService
}

// Get http://localhost:8080/
func (c *AdminResultController) Get() mvc.Result {
	giftID := c.Ctx.URLParamIntDefault("gift_id", 0)
	uid := c.Ctx.URLParamIntDefault("uid", 0)
	page := c.Ctx.URLParamIntDefault("page", 1)

	size := 100
	pagePrev := ""
	pageNext := ""

	var dataList []models.LtResult
	var total int = 0
	// 数据列表
	if giftID > 0 {
		dataList = c.ServiceResult.SearchByGift(giftID, page, size)
	} else if uid > 0 {
		dataList = c.ServiceResult.SearchByUser(uid, page, size)
	} else {
		dataList = c.ServiceResult.GetAll(page, size)
	}

	total = (page - 1) + len(dataList)
	// 数据总数
	if len(dataList) >= size {
		if giftID > 0 {
			total = int(c.ServiceResult.CountByGift(giftID))
		} else if uid > 0 {
			total = int(c.ServiceResult.CountByUser(uid))
		} else {
			total = int(c.ServiceResult.CountAll())
		}
		pageNext = fmt.Sprintf("%d", page+1)
	}
	if page > 1 {
		pagePrev = fmt.Sprintf("%d", page-1)
	}

	log.Println("dataList： ", dataList)
	return mvc.View{
		Name: "admin/result.html",
		Data: iris.Map{
			"Title":    "管理后台",
			"Channel":  "Result",
			"Datalist": dataList,
			"Total":    total,
			"PagePrev": pagePrev,
			"PageNext": pageNext,
			"GiftId":   giftID,
		},
		Layout: "admin/layout.html",
	}
}

// GetCheat http://localhost:8080/cheat
func (c *AdminResultController) GetCheat() mvc.Result {
	giftID, err := c.Ctx.URLParamInt("id")
	if err != nil {
		c.ServiceResult.Update(
			&models.LtResult{Id: giftID, SysStatus: 2},
			[]string{"sys_status"},
		)
	}
	refer := c.Ctx.GetHeader("Referer")
	if refer == "" {
		refer = "/admin/result"
	}
	return mvc.Response{
		Path: refer,
	}
}

// GetDelete http://localhost:8080/delete
func (c *AdminResultController) GetDelete() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	if err == nil {
		c.ServiceResult.Delete(id)
	}
	refer := c.Ctx.GetHeader("Referer")
	if refer == "" {
		refer = "/admin/result"
	}
	return mvc.Response{
		Path: refer,
	}
}

// GetReset http://localhost:8080/reset
func (c *AdminResultController) GetReset() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	if err == nil {
		c.ServiceResult.Update(&models.LtResult{Id: id, SysStatus: 0},
			[]string{"sys_status"})
	}
	refer := c.Ctx.GetHeader("Referer")
	if refer == "" {
		refer = "/admin/result"
	}
	return mvc.Response{
		Path: refer,
	}
}
