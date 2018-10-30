/**
 * 抽奖系统数据处理（包括数据库，也包括缓存等其他形式数据）
 */
package services

import (
				"imooc.com/lottery/dao"
	"imooc.com/lottery/datasource"
	"imooc.com/lottery/models"
		)

type UserService interface {
	GetAll(page, size int) []models.LtUser
	CountAll() int
	//Search(country string) []models.LtUser
	Get(id int) *models.LtUser
	//Delete(id int) error
	Update(user *models.LtUser, columns []string) error
	Create(user *models.LtUser) error
}

type userService struct {
	dao *dao.UserDao
}

func NewUserService() UserService {
	return &userService{
		dao: dao.NewUserDao(datasource.InstanceDbMaster()),
	}
}

func (s *userService) GetAll(page, size int) []models.LtUser {
	return s.dao.GetAll(page, size)
}

func (s *userService) CountAll() int {
	return s.dao.CountAll()
}

//func (s *userService) Search(country string) []models.LtUser {
//	return s.dao.Search(country)
//}

func (s *userService) Get(id int) *models.LtUser {
	data := s.dao.Get(id)
	if data == nil || data.Id <= 0 {
		data = &models.LtUser{Id: id}
	}
	return data
}

//func (s *userService) Delete(id int) error {
//	return s.dao.Delete(id)
//}

func (s *userService) Update(data *models.LtUser, columns []string) error {
	return s.dao.Update(data, columns)
}

func (s *userService) Create(data *models.LtUser) error {
	return s.dao.Create(data)
}
