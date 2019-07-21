package dao

import (
	"log"
	"lottery/models"

	"github.com/go-xorm/xorm"
)

type UserdayDao struct {
	Engine *xorm.Engine
}

func NewUserdayDao(en *xorm.Engine) *UserdayDao {
	return &UserdayDao{
		Engine: en,
	}
}

func (d *UserdayDao) Get(id int) *models.Userday {
	data := &models.Userday{Id: id}
	ok, err := d.Engine.Get(data)
	if !ok || err != nil {
		log.Println("failed to Userday_dao.Get...", err)
		return nil
	}
	return data
}

func (d *UserdayDao) GetAll() []models.Userday {
	dataList := make([]models.Userday, 0)
	err := d.Engine.Asc("sys_created").Find(&dataList)
	if err != nil {
		log.Println("failed to Userday_dao.GetAll...", err)
		return nil
	}
	return dataList
}

func (d *UserdayDao) CountAll() int64 {
	num, err := d.Engine.Count(&models.Userday{})
	if err != nil {
		log.Println("failed to Userday_dao.CountAll...", err)
		return 0
	}
	return num
}

func (d *UserdayDao) Delete(id int) error {
	data := &models.Userday{Id: id}
	_, err := d.Engine.Id(data.Id).Update(data)
	if err != nil {
		log.Println("failed to Userday_dao.Delete...", err)
		return err
	}
	return nil
}

func (d *UserdayDao) Update(data *models.Userday, columns []string) error {
	_, err := d.Engine.Id(data.Id).MustCols(columns...).Update(data)
	if err != nil {
		log.Println("failed to Userday_dao.Update...", err)
		return err
	}
	return nil
}

func (d *UserdayDao) Create(data *models.Userday) error {
	_, err := d.Engine.Insert(data)
	if err != nil {
		log.Println("failed to Userday_dao.Update...", err)
		return err
	}
	return nil
}
