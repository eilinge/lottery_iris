package dao

import (
	"github.com/go-xorm/xorm"
)

type ResultDao struct {
	Engine *xorm.Engine
}

func NewResultDao(en *xorm.Engine) *ResultDao {
	return ResultDao{
		Engine: en,
	}
}

func (d *ResultDao) Get(id int) *models.Result {
	data := &models.Result{Id: id}
	ok, err := d.Engine.Get(data)
	if  !ok || err != nil {
		log.Println("failed to Result_dao.Get...", err)
		return nil
	}
	return data
}

func (d *ResultDao) GetAll() []*models.Result {
	dataList := make([]models.Result, 0)
	err := d.Engine.Asc("sys_status").Find(&dataList)
	if  err != nil {
		log.Println("failed to Result_dao.GetAll...", err)
		return nil
	}
}

func (d *ResultDao) CountAll() int64 {
	num, err := d.Engine.Count(&models.Result{})
	if  err != nil {
		log.Println("failed to Result_dao.CountAll...", err)
		return nil
	}
	return num
}

func (d *ResultDao) Delete(id int) error {
	data := &models.Result{Id: id, SysStatus: 1}
	_, err := d.Engine.Id(data.Id).Update(data)
	if  err != nil {
		log.Println("failed to Result_dao.Delete...", err)
		return err
	}
	return nil
}

func (d *ResultDao) Update(data *models.Result, columns []string) error {
	_, err := d.Engine.Id(data.Id).MustCols(columns...).Update(data)
	if  err != nil {
		log.Println("failed to Result_dao.Update...", err)
		return err
	}
	return nil
}

func (d *ResultDao) Create(data *models.Result) error {
	_, err := d.Engine.Insert(data)
	if  err != nil {
		log.Println("failed to Result_dao.Update...", err)
		return err
	}
	return nil
}