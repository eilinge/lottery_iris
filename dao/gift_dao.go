package dao

import (
	"github.com/go-xorm/xorm"
)

type GiftDao struct {
	Engine *xorm.Engine
}

func NewGiftDao(en *xorm.Engine) *GiftDao {
	return GiftDao{
		Engine: en,
	}
}

func (d *GiftDao) Get(id int) *models.Gift {
	data := &models.Gift{Id: id}
	ok, err := d.Engine.Get(data)
	if  !ok || err != nil {
		log.Println("failed to gift_dao.Get...", err)
		return nil
	}
	return data
}

func (d *GiftDao) GetAll() []*models.Gift {
	dataList := make([]models.Gift, 0)
	err := d.Engine.Asc("sys_status").Asc("displayorder").Find(&dataList)
	if  err != nil {
		log.Println("failed to gift_dao.GetAll...", err)
		return nil
	}
}

func (d *GiftDao) CountAll() int64 {
	num, err := d.Engine.Count(&models.Gift{})
	if  err != nil {
		log.Println("failed to gift_dao.CountAll...", err)
		return nil
	}
	return num
}

func (d *GiftDao) Delete(id int) error {
	data := &models.Gift{Id: id, SysStatus: 1}
	_, err := d.Engine.Id(data.Id).Update(data)
	if  err != nil {
		log.Println("failed to gift_dao.Delete...", err)
		return err
	}
	return nil
}

func (d *GiftDao) Update(data *models.Gift, columns []string) error {
	_, err := d.Engine.Id(data.Id).MustCols(columns...).Update(data)
	if  err != nil {
		log.Println("failed to gift_dao.Update...", err)
		return err
	}
	return nil
}

func (d *GiftDao) Create(data *models.Gift) error {
	_, err := d.Engine.Insert(data)
	if  err != nil {
		log.Println("failed to gift_dao.Update...", err)
		return err
	}
	return nil
}