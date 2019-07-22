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
	"strings"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// AdminCodeController 其他用户访问界面
type AdminCodeController struct {
	Ctx         iris.Context // 解析前端传来的数据
	ServiceCode services.CodeService
	ServiceGift services.GiftService
}

// Get http://localhost:8080/
func (c *AdminCodeController) Get() mvc.Result {
	giftID := c.Ctx.URLParamIntDefault("gift_id", 0)
	page := c.Ctx.URLParamIntDefault("page", 1)

	size := 100
	pagePrev := ""
	pageNext := ""

	var dataList []models.LtCode
	var total int = 0

	if giftID > 0 {
		dataList = c.ServiceCode.Search(giftID)
	} else {
		dataList = c.ServiceCode.GetAll(page, size)
	}

	total = (page - 1) + len(dataList)
	if len(dataList) >= size {
		if giftID > 0 {
			total = int(c.ServiceCode.CountByGift(giftID))
		} else {
			total = int(c.ServiceCode.CountAll())
		}
		pageNext = fmt.Sprintf("%d", page+1)
	}
	if page > 1 {
		pagePrev = fmt.Sprintf("%d", page-1)
	}

	log.Println("dataList： ", dataList)
	return mvc.View{
		Name: "admin/code.html",
		Data: iris.Map{
			"Title":    "管理后台",
			"Channel":  "Code",
			"Datalist": dataList,
			"Total":    total,
			"PagePrev": pagePrev,
			"PageNext": pageNext,
			"GiftId":   giftID,
		},
		Layout: "admin/layout.html",
	}
}

// PostImport http://localhost:8080/import
func (c *AdminCodeController) PostImport() {
	giftID := c.Ctx.URLParamIntDefault("gift_id", 0)
	if giftID < 1 {
		c.Ctx.HTML("没有指定奖品ID, 无法进行导入, <a href='' onclick='history.go(-1); return false;'>返回</a>")
		return
	}
	gift := c.ServiceGift.Get(giftID, false)
	// if gift == nil || gift.Id < 1 || gift.Gtype != conf.GtypeCodeDiff { // ?
	if gift == nil || gift.Id < 1 { // ?
		c.Ctx.HTML("奖品信息不存在或奖品类型不是差异化优惠券, 无法进行导入, <a href='' onclick='history.go(-1); return false;'>返回</a>")
		return
	}

	codes := c.Ctx.PostValue("codes")
	now := comm.NowUnix()
	list := strings.Split(codes, "\n")
	sucNum := 0
	errNum := 0

	for _, code := range list {
		code = strings.TrimSpace(code)
		if code != "" {
			data := &models.LtCode{
				GiftId:     giftID,
				Code:       code,
				SysCreated: now,
			}
			err := c.ServiceCode.Create(data)
			log.Println("err: ", err)
			if err != nil {
				errNum++
			} else {
				sucNum++
				// TODO: 成功导入数据库, 下一步还需要导入缓存
			}
		}
	}
	c.Ctx.HTML(fmt.Sprintf("成功导入%d条, 导入失败%d条, <a href='/admin/code?gift_id=%d'>返回</a>", sucNum, errNum, giftID))
}

// GetDelete http://localhost:8080/delete
func (c *AdminCodeController) GetDelete() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	if err == nil {
		c.ServiceCode.Delete(id)
	}
	refer := c.Ctx.GetHeader("Referer")
	if refer == "" {
		refer = "/admin/code"
	}
	return mvc.Response{
		Path: refer,
	}
}

// GetReset http://localhost:8080/reset
func (c *AdminCodeController) GetReset() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	if err == nil {
		c.ServiceCode.Update(&models.LtCode{Id: id, SysStatus: 0},
			[]string{"sys_status"})
	}
	refer := c.Ctx.GetHeader("Referer")
	if refer == "" {
		refer = "/admin/code"
	}
	return mvc.Response{
		Path: refer,
	}
}
