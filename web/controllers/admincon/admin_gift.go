/**
 * 首页根目录的Controller
 * http://localhost:8080/
 */
package admincon

import (
	"fmt"
	"encoding/json"
	"time"
	
	"lottery/services"
	"lottery/comm"
	"lottery/models"
	"lottery/web/viewmodels"
	"lottery/web/utils"
	
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// AdminGiftController 其他用户访问界面
type AdminGiftController struct {
	Ctx            iris.Context // 解析前端传来的数据
	ServiceGift    services.GiftService
}

// Get http://localhost:8080/
func (c *AdminGiftController) Get() mvc.Result {
	dataList := c.ServiceGift.GetAll(false)
	total := len(dataList)

	for i, giftInfo := range dataList {
		prizedata := make([][2]int, 0)
		err := json.Unmarshal([]byte(giftInfo.PrizeData), &prizedata)
		if err != nil || len(prizedata) < 1{
			dataList[i].PrizeData = "[]"
		} else {
			// timestamp type int -> string
			newpd := make([]string, len(prizedata))
			for index, pd := range prizedata {
				ct := comm.FormatFromUnixTime(int64(pd[0]))
				newpd[index] = fmt.Sprintf("[%s]: [%d]", ct, pd[1])
			}
			str, err := json.Marshal(newpd)
			if err == nil && len(str) >0 {
				dataList[i].PrizeData = string(str)
			} else {
				dataList[i].PrizeData = "[]"
			}
		}
		num := utils.GetGiftPoolNum(giftInfo.Id)
		dataList[i].Title = fmt.Sprintf("[%d] %s", num, dataList[i].Title)
	}

	return mvc.View{
		Name: "admin/gift.html",
		Data: iris.Map{
			"Title": "管理后台",
			"Channel": "gift",
			"Datalist": dataList,
			"Total": total,
		},
		Layout: "admin/layout.html",
	}
}

// GetEdit http://localhost:8080/edit
func (c *AdminGiftController) GetEdit() mvc.Result {
	id := c.Ctx.URLParamIntDefault("id", 0)
	
	giftInfo := viewmodels.ViewGift{}
	if id > 0 {
		data := c.ServiceGift.Get(id, false)
		fmt.Println("data: ", data)
		giftInfo.Id = data.Id
		giftInfo.Title = data.Title
		giftInfo.PrizeNum = data.PrizeNum
		giftInfo.PrizeCode = data.PrizeCode
		giftInfo.PrizeTime = data.PrizeTime
		giftInfo.Img = data.Img
		giftInfo.Displayorder = data.Displayorder
		giftInfo.Gtype = data.Gtype
		giftInfo.Gdata = data.Gdata
		giftInfo.TimeBegin = comm.FormatFromUnixTime(int64(data.TimeBegin))
		giftInfo.TimeEnd = comm.FormatFromUnixTime(int64(data.TimeEnd))
	}
	return mvc.View{
		Name: "admin/giftEdit.html",
		Data: iris.Map{
			"Title": "管理后台",
			"Channel": "gift",
			"info": giftInfo,
		},
		Layout: "admin/layout.html",
	}
}

// PostSave http://localhost:8080/save
func (c *AdminGiftController) PostSave() mvc.Result {
	data := viewmodels.ViewGift{}
	err := c.Ctx.ReadForm(&data)
	if err != nil {
		fmt.Println("admin_gift.PostSave ReadForm error=", err)
		return mvc.Response{
			Text: fmt.Sprintf("ReadForm 转换异常, err=%s", err),
		}
	}
	giftInfo := models.LtGift{}
	giftInfo.Id = data.Id
	giftInfo.Title = data.Title
	giftInfo.PrizeNum = data.PrizeNum
	giftInfo.PrizeCode = data.PrizeCode
	giftInfo.PrizeTime = data.PrizeTime
	giftInfo.Img = data.Img
	giftInfo.Displayorder = data.Displayorder
	giftInfo.Gtype = data.Gtype
	giftInfo.Gdata = data.Gdata

	t1, err1 := comm.ParseTime(data.TimeBegin)
	t2, err2 := comm.ParseTime(data.TimeEnd)
	if err1 != nil || err2 != nil {
		return mvc.Response{
			Text: fmt.Sprintf("开始时间, 结束时间不正确, err1=%s, err2=%s", err1, err2),
		}
	}
	giftInfo.TimeBegin = int(t1.Unix())
	giftInfo.TimeEnd = int(t2.Unix())

	if giftInfo.Id > 0 {
		// 数据更新
		datainfo := c.ServiceGift.Get(giftInfo.Id, false) // 数据库
		if datainfo != nil && datainfo.Id > 0 {
			// giftinfo 表单中数据
			if datainfo.PrizeNum != giftInfo.PrizeNum {
				// 剩余数 - 总数的变化
				giftInfo.LeftNum = datainfo.LeftNum - (datainfo.PrizeNum - giftInfo.PrizeNum)
				if giftInfo.LeftNum < 0 || giftInfo.PrizeNum <= 0 {
					giftInfo.LeftNum = 0
				}
				// 奖品总数发生了变化
				utils.ResetGiftPrizeData(&giftInfo, c.ServiceGift)
			}
			if datainfo.PrizeTime != giftInfo.PrizeTime {
				// 发奖周期发生了变化
				utils.ResetGiftPrizeData(&giftInfo, c.ServiceGift)
			}
			giftInfo.SysUpdated = int(time.Now().Unix())
			c.ServiceGift.Update(&giftInfo, []string{""})
		} else {
			giftInfo.Id = 0
		}
	
	}
	if giftInfo.Id == 0{
		// 数据添加
		giftInfo.LeftNum = giftInfo.PrizeNum
		giftInfo.SysIp = comm.ClientIP(c.Ctx.Request())
		giftInfo.SysCreated = int(time.Now().Unix())
		c.ServiceGift.Create(&giftInfo)
		// 新的奖品, 更新奖品的发奖计划
		utils.ResetGiftPrizeData(&giftInfo, c.ServiceGift)
	}
	return mvc.Response{
		Path: "/admin/gift",
	}
}

// GetDelete http://localhost:8080/Delete
func (c *AdminGiftController) GetDelete() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	if err == nil {
		c.ServiceGift.Delete(id)
	}
	return mvc.Response{
		Path: "/admin/gift",
	}
}

// GetReset http://localhost:8080/reset
func (c *AdminGiftController) GetReset() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	if err == nil {
		c.ServiceGift.Update(&models.LtGift{Id: id, SysStatus:0},
			[]string{"sys_status"})
	}
	return mvc.Response{
		Path: "/admin/gift",
	}
}