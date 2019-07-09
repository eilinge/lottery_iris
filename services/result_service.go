package services

import (
	"lottery/dao"
	"lottery/models"
)

type ResultService interface {
	GetAll() []*modles.Result
	CountAll() int64
	Get(id int) *models.Result
	Delete(id int) error
	Update(date *models.Result, columns []string) error
	Create(data *models.Result) error
}

type resultService struct {
	Dao *dao.ResultDao
}

func NewresultService() *ResultService {
	return &resultService{
		Dao: dao.NewResultDao(nil),
	}
}

func (s *resultService) GetAll() []*modles.Result {
	return s.Dao.GetAll()
}

func (s *resultService) CountAll() int64 {
	return s.Dao.CountAll()
}

func (s *resultService) Get(id int) *models.Result {
	return s.Dao.Get(id)
}

func (s *resultService) Delete(id int) error {
	return s.Dao.Delete(id)
}

func (s *resultService) Update(date *models.Result, columns []string) error {
	return s.Dao.Update(data, columns)
}

func (s *resultService) Create(data *models.Result) error {
	return s.Dao.Create(data)
}