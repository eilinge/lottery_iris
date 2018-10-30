/**
 * 抽奖系统数据处理（包括数据库，也包括缓存等其他形式数据）
 */
package services

import (
			"strconv"
	"strings"

		"imooc.com/lottery/dao"
	"imooc.com/lottery/datasource"
	"imooc.com/lottery/models"
)

type GiftService interface {
	GetAll() []models.LtGift
	CountAll() int64
	//Search(country string) []models.LtGift
	Get(id int) *models.LtGift
	Delete(id int) error
	Update(data *models.LtGift, columns []string) error
	Create(data *models.LtGift) error
	GetAllUse(useCache bool) []models.ObjGiftPrize
	IncrLeftNum(id, num int) (int64, error)
	DecrLeftNum(id, num int) (int64, error)
}

type giftService struct {
	dao *dao.GiftDao
}

func NewGiftService() GiftService {
	return &giftService{
		dao: dao.NewGiftDao(datasource.InstanceDbMaster()),
	}
}

func (s *giftService) GetAll() []models.LtGift {
	return s.dao.GetAll()
}

func (s *giftService) CountAll() int64 {
	return s.dao.CountAll()
}

//func (s *giftService) Search(country string) []models.LtGift {
//	return s.dao.Search(country)
//}

func (s *giftService) Get(id int) *models.LtGift {
	return s.dao.Get(id)
}

func (s *giftService) Delete(id int) error {
	return s.dao.Delete(id)
}

func (s *giftService) Update(data *models.LtGift, columns []string) error {
	return s.dao.Update(data, columns)
}

func (s *giftService) Create(data *models.LtGift) error {
	return s.dao.Create(data)
}

// 获取到当前可以获取的奖品列表
// 有奖品限定，状态正常，时间期间内
// gtype倒序， displayorder正序
func (s *giftService) GetAllUse(useCache bool) []models.ObjGiftPrize {
	list := make([]models.LtGift, 0)
	// 直接读取数据库的方式
	list = s.dao.GetAllUse()

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
	} else {
		return []models.ObjGiftPrize{}
	}
}

func (s *giftService) IncrLeftNum(id, num int) (int64, error) {
	return s.dao.IncrLeftNum(id, num)
}

func (s *giftService) DecrLeftNum(id, num int) (int64, error) {
	return s.dao.DecrLeftNum(id, num)
}
