package dao

import (
	"log"
	"lottery/models"

	"github.com/go-xorm/xorm"
)

type BlackipDao struct {
	Engine *xorm.Engine
}

func NewBlackipDao(en *xorm.Engine) *BlackipDao {
	return &BlackipDao{
		Engine: en,
	}
}

func (d *BlackipDao) Get(id int) *models.Blackip {
	data := &models.Blackip{Id: id}
	ok, err := d.Engine.Get(data)
	if !ok || err != nil {
		log.Println("failed to Blackip_dao.Get...", err)
		return nil
	}
	return data
}

func (d *BlackipDao) GetAll() []models.Blackip {
	dataList := make([]models.Blackip, 0)
	err := d.Engine.Asc("sys_status").Asc("displayorder").Find(&dataList)
	if err != nil {
		log.Println("failed to Blackip_dao.GetAll...", err)
		return nil
	}
	return dataList
}

func (d *BlackipDao) CountAll() int64 {
	num, err := d.Engine.Count(&models.Blackip{})
	if err != nil {
		log.Println("failed to Blackip_dao.CountAll...", err)
		return 0
	}
	return num
}

func (d *BlackipDao) Delete(id int) error {
	data := &models.Blackip{Id: id}
	_, err := d.Engine.Id(data.Id).Update(data)
	if err != nil {
		log.Println("failed to Blackip_dao.Delete...", err)
		return err
	}
	return nil
}

func (d *BlackipDao) Update(data *models.Blackip, columns []string) error {
	_, err := d.Engine.Id(data.Id).MustCols(columns...).Update(data)
	if err != nil {
		log.Println("failed to Blackip_dao.Update...", err)
		return err
	}
	return nil
}

func (d *BlackipDao) Create(data *models.Blackip) error {
	_, err := d.Engine.Insert(data)
	if err != nil {
		log.Println("failed to Blackip_dao.Update...", err)
		return err
	}
	return nil
}

func (d *BlackipDao) GetById(ip string) *models.Blackip {
	dataList := make([]models.Blackip, 0)
	err := d.Engine.Where("ip=?", ip).Desc("id").Limit(1).Find(&dataList)
	if err != nil || len(dataList) < 1 {
		return nil
	} else {
		return &dataList[0]
	}
}
