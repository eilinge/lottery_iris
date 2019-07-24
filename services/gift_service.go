/**
 * 抽奖系统数据处理（包括数据库，也包括缓存等其他形式数据）
 */
package services

import (
	"encoding/json"
	"time"
	"log"
	"strconv"
	"strings"

	"lottery/comm"
	"lottery/dao"
	"lottery/datasource"
	"lottery/models"
)

type GiftService interface {
	GetAll(useCache bool) []models.LtGift
	CountAll() int64
	//Search(country string) []models.LtGift
	Get(id int, useCache bool) *models.LtGift
	Delete(id int) error
	Update(data *models.LtGift, columns []string) error
	Create(data *models.LtGift) error
	GetAllUse(useCache bool) []models.ObjGiftPrize
	IncrLeftNum(id, num int) (int64, error)
	DecrLeftNum(id, num int) (int64, error)
	getAllByCache() []models.LtGift
	setAllByCache(gifts []models.LtGift)
	updateByCache(data *models.LtGift, columns []string)
}

type giftService struct {
	dao *dao.GiftDao
}

// NewGiftService entrance of gitservice
func NewGiftService() GiftService {
	return &giftService{
		dao: dao.NewGiftDao(datasource.InstanceDbMaster()),
	}
}

func (s *giftService) GetAll(useCache bool) []models.LtGift {
	if !useCache {
		// 直接读取数据库的方式
		return s.dao.GetAll()
	}
	// 奖品的更新频率较低, 90%会直接读取缓存
	gifts := s.getAllByCache()
	if len(gifts) < 1 {
		gifts = s.dao.GetAll()
		s.setAllByCache(gifts)
	}	

	return gifts
}

func (s *giftService) CountAll() int64 {
	// 直接读取数据库的方式
	//return s.dao.CountAll()

	// 缓存优化之后的读取方式
	gifts := s.GetAll(true)
	return int64(len(gifts))
}

func (s *giftService) Get(id int, useCache bool) *models.LtGift {
	if !useCache {
		// 直接读取数据库的方式
		return s.dao.Get(id)
	}
	gifts := s.GetAll(true)
	for _, gift := range gifts {
		if gift.Id == id {
			return &gift
		}
	}
	return nil
}

// 更新数据时(delete, update, create), 必须先更新缓存

func (s *giftService) Delete(id int) error {
	data := &models.LtGift{Id: id}
	s.updateByCache(data, nil)
	// 再更新数据库
	return s.dao.Delete(id)
}

func (s *giftService) Update(data *models.LtGift, columns []string) error {
	s.updateByCache(data, columns)
	// 再更新数据库
	return s.dao.Update(data, columns)
}

func (s *giftService) Create(data *models.LtGift) error {
	s.updateByCache(data, nil)
	// 再更新数据库
	return s.dao.Create(data)
}

// GetAllUse ...
// 获取到当前可以获取的奖品列表
// 有奖品限定，状态正常，时间期间内
// gtype倒序， displayorder正序
func (s *giftService) GetAllUse(useCache bool) []models.ObjGiftPrize {
	list := make([]models.LtGift, 0)
	
	if !useCache {
		list = s.dao.GetAllUse()
	} else {
		now := time.Now().Unix()
		gifts := s.GetAll(true)
		for _, gift := range gifts {
			if gift.Id > 0 && gift.SysStatus == 0 && gift.PrizeNum >= 0 && gift.TimeBegin <= int(now) && gift.TimeEnd >= int(now) {
				list = append(list, gift)
			}
		}
	}
	if list != nil {
		gifts := make([]models.ObjGiftPrize, 0)
		for _, gift := range list {
			codes := strings.Split(gift.PrizeCode, "-")
			if len(codes) == 2 {
				// 设置了获奖编码范围 a-b 才可以进行抽奖
				codeA := codes[0]
				codeB := codes[1]
				a, e1 := strconv.Atoi(codeA)
				b, e2 := strconv.Atoi(codeB)
				if e1 == nil && e2 == nil && b >= a && a >= 0 && b < 10000 {
					data := models.ObjGiftPrize{
						Id:           gift.Id,
						Title:        gift.Title,
						PrizeNum:     gift.PrizeNum,
						LeftNum:      gift.LeftNum,
						PrizeCodeA:   a,
						PrizeCodeB:   b,
						Img:          gift.Img,
						Displayorder: gift.Displayorder,
						Gtype:        gift.Gtype,
						Gdata:        gift.Gdata,
					}
					gifts = append(gifts, data)
				}
			}
		}
		return gifts
	}
	return []models.ObjGiftPrize{}

}

func (s *giftService) IncrLeftNum(id, num int) (int64, error) {
	return s.dao.IncrLeftNum(id, num)
}

func (s *giftService) DecrLeftNum(id, num int) (int64, error) {
	return s.dao.DecrLeftNum(id, num)
}

/*
奖品数据全量缓存, Json序列化为string结构, gift_service
string(简单): 后台的奖品更新频率较低, JSON序列化, 直接保存到string中 
hash(无序): 奖品数据需要排序
list(分别对多个奖品数据进行管理): 工作量太大, 体现不到list的优势

1. 增加了3个方法: getAllByCache, setAllByCache, updateByCache
2. 修改读取方法, 增加userCache bool 参数
3. 修改数据的时候, 清空缓存, 下次读取的时候自动更新最新数据
*/

func (s *giftService) getAllByCache() []models.LtGift {
	key := "allgift"
	rds := datasource.InstanceCache()
	rs, err := rds.Do("GET", key)

	if err != nil {
		log.Println("gift_service.getAllByCache Get key= ", key, ", error = ", err)
		return nil
	}
	str := comm.GetString(rs, "")
	if str == "" {
		return nil
	}
	// LtGift中设置了json:"-", 无法直接用[]models.LtGift做序列化和反序列化处理
	dataList := []map[string]interface{}{}
	err = json.Unmarshal([]byte(str), &dataList)
	if err != nil {
		log.Println("gift_service.getAllByCache json.Unmarshal error = ", err)
		return nil
	}
	gifts := make([]models.LtGift, len(dataList))
	for i := 0; i<len(dataList); i++ {
		data := dataList[i]
		id := comm.GetInt64FromMap(data, "Id", 0)
		if id < 0 {
			gifts[i] = models.LtGift{}
		} else {
			gift := models.LtGift{
				Id: int(id),
				Title: comm.GetStringFromMap(data, "Title", ""),
				PrizeNum: int(comm.GetInt64FromMap(data, "PrizeNum", 0)),
				LeftNum: int(comm.GetInt64FromMap(data, "LeftNum", 0)),
				PrizeCode: comm.GetStringFromMap(data, "PrizeCode", ""),
				PrizeTime: int(comm.GetInt64FromMap(data, "PrizeTime", 0)),
				Img: comm.GetStringFromMap(data, "Img", ""),
				Displayorder: int(comm.GetInt64FromMap(data, "Displayorder", 0)),
				Gtype: int(comm.GetInt64FromMap(data, "Gtype", 0)),
				Gdata: comm.GetStringFromMap(data, "Gdata", ""),
				TimeBegin: int(comm.GetInt64FromMap(data, "TimeBegin", 0)),
				TimeEnd: int(comm.GetInt64FromMap(data, "TimeEnd", 0)),
				// PrizeData: comm.GetStringFromMap(data, "PrizeData", ""),
				PrizeBegin: int(comm.GetInt64FromMap(data, "PrizeBegin", 0)),
				PrizeEnd: int(comm.GetInt64FromMap(data, "PrizeEnd", 0)),
				SysStatus: int(comm.GetInt64FromMap(data, "SysStatus", 0)),
				SysCreated: int(comm.GetInt64FromMap(data, "SysCreated", 0)),
				SysIp: comm.GetStringFromMap(data, "SysIp", ""),
			}
			gifts[i] = gift
		}
	}
	return gifts
}

func (s *giftService) setAllByCache(gifts []models.LtGift) {
	strValue := ""
	if len(gifts) > 0 {
		dataList := make([]map[string]interface{}, len(gifts))
		for i := 0; i < len(dataList); i++ {
			gift := gifts[i]
			data := make(map[string]interface{})
			data["Id"] = gift.Id
			data["Title"] = gift.Title
			data["PrizeNum"] = gift.PrizeNum
			data["LeftNum"] = gift.LeftNum
			data["PrizeCode"] = gift.PrizeCode
			data["PrizeTime"] = gift.PrizeTime
			data["Img"] = gift.Img
			data["Displayorder"] = gift.Displayorder
			data["Gtype"] = gift.Gtype
			data["Gdata"] = gift.Gdata
			data["TimeBegin"] = gift.TimeBegin
			data["TimeEnd"] = gift.TimeEnd
			data["PrizeBegin"] = gift.PrizeBegin
			data["PrizeEnd"] = gift.PrizeEnd
			data["SysStatus"] = gift.SysStatus
			data["SysCreated"] = gift.SysCreated
			data["SysIp"] = gift.SysIp
			dataList[i] = data
		}
		str, err := json.Marshal(dataList)
		if err != nil {
			log.Println("gift_service.setAllByCache json.Marshal error = ", err)
		}
		strValue = string(str)
	}
	key := "allgift"
	rds := datasource.InstanceCache()
	_, err := rds.Do("SET", key, strValue)
	if err != nil {
		log.Println("gift_service.setAllByCache redis SET key=", key, ", error = ", err)
	}
}

// updateByCache 在缓存中对key相应的value进行更新会比较繁琐, 先删除对应key, 再从数据库中读取
func (s *giftService) updateByCache(data *models.LtGift, columns []string) {
	if data == nil || data.Id <= 0 {
		return
	}
	key := "allgift"
	rds := datasource.InstanceCache()
	_, err := rds.Do("DEL", key)
	if err != nil {
		log.Println("gift_service.updateByCache redis DEL key=", key, ", error = ", err)
	}
}