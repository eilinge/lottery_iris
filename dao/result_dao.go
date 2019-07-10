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

func (d *ResultDao) GetAll(page, size int) []*models.Result {
	offset := (page - 1) * size
	dataList := make([]models.Result, 0)
	err := d.Engine.Desc("id").Limit(size, offset).Find(&dataList)
	if  err != nil {
		log.Println("failed to Result_dao.GetAll...", err)
		return nil
	}
	return dataList
}

func (d *ResultDao) CountAll() int64 {
	num, err := d.Engine.Count(&models.Result{})
	if  err != nil {
		log.Println("failed to Result_dao.CountAll...", err)
		return nil
	}
	return num
}

func (d *ResultDao) GetNewPrize(size int, giftIds []int) []models.Result{
	dataList := make([]models.Result, 0)
	err := d.Engine.In("gift_id", giftIds).Desc("id").Limit(size).Find(&datalist)
	if  err != nil {
		log.Println("failed to Result_dao.GetNewPrize...", err)
		return nil
	}
	return dataList
}

func (d *ResultDao) SearchByGift(giftId, page, size int) []models.Result{
	offset := (page - 1) * size
	dataList := make([]models.Result, 0)
	err := d.Engine.Where("gift_id=?", giftIds).Desc("id").Limit(size, offset).Find(&datalist)
	if  err != nil {
		log.Println("failed to Result_dao.SearchByGift...", err)
		return nil
	}
	return dataList
}

func (d *ResultDao) CountByGift(giftId int) int64 {
	
	num, err := d.Engine.Where("gift_id=?", giftId).Count(&models.Result{})
	if  err != nil {
		log.Println("failed to Result_dao.CountByGift...", err)
		return nil
	}
	retrun num
}

func (d *ResultDao) CountByUser(uid int) int64 {
	
	num, err := d.Engine.Where("uid=?", uid).Count(&models.Result{})
	if  err != nil {
		log.Println("failed to Result_dao.CountByUser...", err)
		return 0
	}
	retrun num
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
		log.Println("failed to Result_dao.Create...", err)
		return err
	}
	return nil
}