package dao

import (
	"github.com/go-xorm/xorm"
)

type CodeDao struct {
	Engine *xorm.Engine
}

func NewCodeDao(en *xorm.Engine) *CodeDao {
	return CodeDao{
		Engine: en,
	}
}

func (d *CodeDao) Get(id int) *models.Code {
	data := &models.Code{Id: id}
	ok, err := d.Engine.Get(data)
	if  !ok || err != nil {
		log.Println("failed to Code_dao.Get...", err)
		return nil
	}
	return data
}

func (d *CodeDao) GetAll() []*models.Code {
	dataList := make([]models.Code, 0)
	err := d.Engine.Desc("id").Find(&dataList)
	if  err != nil {
		log.Println("failed to Code_dao.GetAll...", err)
		return nil
	}
}

func (d *CodeDao) CountAll() int64 {
	num, err := d.Engine.Count(&models.Code{})
	if  err != nil {
		log.Println("failed to Code_dao.CountAll...", err)
		return nil
	}
	return num
}

func (d *CodeDao) Delete(id int) error {
	data := &models.Code{Id: id, SysStatus: 1}
	_, err := d.Engine.Id(data.Id).Update(data)
	if  err != nil {
		log.Println("failed to Code_dao.Delete...", err)
		return err
	}
	return nil
}

func (d *CodeDao) Update(data *models.Code, columns []string) error {
	_, err := d.Engine.Id(data.Id).MustCols(columns...).Update(data)
	if  err != nil {
		log.Println("failed to Code_dao.Update...", err)
		return err
	}
	return nil
}

func (d *CodeDao) Create(data *models.Code) error {
	_, err := d.Engine.Insert(data)
	if  err != nil {
		log.Println("failed to Code_dao.Update...", err)
		return err
	}
	return nil
}