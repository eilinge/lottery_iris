package services

import (
	"lottery/dao"
	"lottery/models"
)

type CodeService interface {
	GetAll() []models.Code
	CountAll() int64
	Get(id int) *models.Code
	Delete(id int) error
	Update(date *models.Code, columns []string) error
	Create(data *models.Code) error
}

type codeService struct {
	Dao *dao.CodeDao
}

func NewCodeService() *codeService {
	return &codeService{
		Dao: dao.NewCodeDao(nil),
	}
}

func (s *codeService) GetAll() []*models.Code {
	return s.Dao.GetAll()
}

func (s *codeService) CountAll() int64 {
	return s.Dao.CountAll()
}

func (s *codeService) Get(id int) *models.Code {
	return s.Dao.Get(id)
}

func (s *codeService) Delete(id int) error {
	return s.Dao.Delete(id)
}

func (s *codeService) Update(data *models.Code, columns []string) error {
	return s.Dao.Update(data, columns)
}

func (s *codeService) Create(data *models.Code) error {
	return s.Dao.Create(data)
}
