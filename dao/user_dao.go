package dao

import (
	"github.com/go-xorm/xorm"
)

type UserDao struct {
	Engine *xorm.Engine
}

func NewUserDao(en *xorm.Engine) *UserDao {
	return UserDao{
		Engine: en,
	}
}

func (d *UserDao) Get(id int) *models.User {
	data := &models.User{Id: id}
	ok, err := d.Engine.Get(data)
	if  !ok || err != nil {
		log.Println("failed to User_dao.Get...", err)
		return nil
	}
	return data
}

func (d *UserDao) GetAll() []*models.User {
	dataList := make([]models.User, 0)
	err := d.Engine.Find(&dataList)
	if  err != nil {
		log.Println("failed to User_dao.GetAll...", err)
		return nil
	}
}

func (d *UserDao) CountAll() int64 {
	num, err := d.Engine.Count(&models.User{})
	if  err != nil {
		log.Println("failed to User_dao.CountAll...", err)
		return nil
	}
	return num
}

func (d *UserDao) Delete(id int) error {
	data := &models.User{Id: id, SysStatus: 1}
	_, err := d.Engine.Id(data.Id).Update(data)
	if  err != nil {
		log.Println("failed to User_dao.Delete...", err)
		return err
	}
	return nil
}

func (d *UserDao) Update(data *models.User, columns []string) error {
	_, err := d.Engine.Id(data.Id).MustCols(columns...).Update(data)
	if  err != nil {
		log.Println("failed to User_dao.Update...", err)
		return err
	}
	return nil
}

func (d *UserDao) Create(data *models.User) error {
	_, err := d.Engine.Insert(data)
	if  err != nil {
		log.Println("failed to User_dao.Update...", err)
		return err
	}
	return nil
}