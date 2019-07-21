package services

import (
	"lottery/dao"
	"lottery/models"
)

type UserdayService interface {
	GetAll() []models.Userday
	CountAll() int64
	Get(id int) *models.Userday
	Delete(id int) error
	Update(date *models.Userday, columns []string) error
	Create(data *models.Userday) error
}

type userdayService struct {
	Dao *dao.UserdayDao
}

func NewUserdayService() *userdayService {
	return &userdayService{
		Dao: dao.NewUserdayDao(nil),
	}
}

func (s *userdayService) GetAll() []models.Userday {
	return s.Dao.GetAll()
}

func (s *userdayService) CountAll() int64 {
	return s.Dao.CountAll()
}

func (s *userdayService) Get(id int) *models.Userday {
	return s.Dao.Get(id)
}

func (s *userdayService) Delete(id int) error {
	return s.Dao.Delete(id)
}

func (s *userdayService) Update(data *models.Userday, columns []string) error {
	return s.Dao.Update(data, columns)
}

func (s *userdayService) Create(data *models.Userday) error {
	return s.Dao.Create(data)
}
