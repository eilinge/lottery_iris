package services

import (
	"lottery/dao"
	"lottery/models"
)

type blackipService interface {
	GetAll() []*modles.Blackip
	CountAll() int64
	Get(id int) *models.Blackip
	Delete(id int) error
	Update(date *models.Blackip, columns []string) error
	Create(data *models.Blackip) error
	GetById(ip string) *models.Blackip
}

type blackipService struct {
	Dao *dao.BlackipDao
}

func NewblackipService() *BlackipService {
	return &blackipService{
		Dao: dao.NewBlackipDao(nil),
	}
}

func (s *blackipService) GetAll() []*modles.Blackip {
	return s.Dao.GetAll()
}

func (s *blackipService) CountAll() int64 {
	return s.Dao.CountAll()
}

func (s *blackipService) Get(id int) *models.Blackip {
	return s.Dao.Get(id)
}

func (s *blackipService) Delete(id int) error {
	return s.Dao.Delete(id)
}

func (s *blackipService) Update(date *models.Blackip, columns []string) error {
	return s.Dao.Update(data, columns)
}

func (s *blackipService) Create(data *models.Blackip) error {
	return s.Dao.Create(data)
}

func (s *blackipService) GetById(ip string) *models.Blackip {
	return s.Dao.GetById(ip)
}