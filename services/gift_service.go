package services

import (
	"lottery/dao"
	"lottery/models"
)

type GiftService interface {
	GetAll() []models.Gift
	CountAll() int64
	Get(id int) *models.Gift
	Delete(id int) error
	Update(date *models.Gift, columns []string) error
	Create(data *models.Gift) error
}

type giftService struct {
	Dao *dao.GiftDao
}

func NewGiftService() *giftService {
	return &giftService{
		Dao: dao.NewGiftDao(nil),
	}
}

func (s *giftService) GetAll() []models.Gift {
	return s.Dao.GetAll()
}

func (s *giftService) CountAll() int64 {
	return s.Dao.CountAll()
}

func (s *giftService) Get(id int) *models.Gift {
	return s.Dao.Get(id)
}

func (s *giftService) Delete(id int) error {
	return s.Dao.Delete(id)
}

func (s *giftService) Update(data *models.Gift, columns []string) error {
	return s.Dao.Update(data, columns)
}

func (s *giftService) Create(data *models.Gift) error {
	return s.Dao.Create(data)
}
