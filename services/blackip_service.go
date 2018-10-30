/**
 * 抽奖系统数据处理（包括数据库，也包括缓存等其他形式数据）
 */
package services

import (
	"imooc.com/lottery/dao"
	"imooc.com/lottery/datasource"
	"imooc.com/lottery/models"
						)

type BlackipService interface {
	GetAll(page, size int) []models.LtBlackip
	CountAll() int64
	Search(ip string) []models.LtBlackip
	Get(id int) *models.LtBlackip
	//Delete(id int) error
	Update(user *models.LtBlackip, columns []string) error
	Create(user *models.LtBlackip) error
	GetByIp(ip string) *models.LtBlackip
}

type blackipService struct {
	dao *dao.BlackipDao
}

func NewBlackipService() BlackipService {
	return &blackipService{
		dao: dao.NewBlackipDao(datasource.InstanceDbMaster()),
	}
}

func (s *blackipService) GetAll(page, size int) []models.LtBlackip {
	return s.dao.GetAll(page, size)
}

func (s *blackipService) CountAll() int64 {
	return s.dao.CountAll()
}

func (s *blackipService) Search(ip string) []models.LtBlackip {
	return s.dao.Search(ip)
}

func (s *blackipService) Get(id int) *models.LtBlackip {
	return s.dao.Get(id)
}

//func (s *blackipService) Delete(id int) error {
//	return s.dao.Delete(id)
//}

func (s *blackipService) Update(data *models.LtBlackip, columns []string) error {
	return s.dao.Update(data, columns)
}

func (s *blackipService) Create(data *models.LtBlackip) error {
	return s.dao.Create(data)
}

// 根据IP读取IP的黑名单数据
func (s *blackipService) GetByIp(ip string) *models.LtBlackip {
	data := s.dao.GetByIp(ip)
	if data == nil || data.Ip == "" {
		data = &models.LtBlackip{Ip: ip}
	}
	return data
}