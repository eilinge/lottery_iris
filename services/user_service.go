package services

import (
	"lottery/dao"
	"lottery/models"
)

type UserService interface {
	GetAll() []*modles.User
	CountAll() int64
	Get(id int) *models.User
	Delete(id int) error
	Update(date *models.User, columns []string) error
	Create(data *models.User) error
}

type userService struct {
	Dao *dao.UserDao
}

func NewuserService() *UserService {
	return &userService{
		Dao: dao.NewUserDao(nil),
	}
}

func (s *userService) GetAll() []*modles.User {
	return s.Dao.GetAll()
}

func (s *userService) CountAll() int64 {
	return s.Dao.CountAll()
}

func (s *userService) Get(id int) *models.User {
	return s.Dao.Get(id)
}

func (s *userService) Delete(id int) error {
	return s.Dao.Delete(id)
}

func (s *userService) Update(date *models.User, columns []string) error {
	return s.Dao.Update(data, columns)
}

func (s *userService) Create(data *models.User) error {
	return s.Dao.Create(data)
}