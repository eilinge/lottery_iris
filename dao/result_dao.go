package dao

import (
	"github.com/go-xorm/xorm"

	"lottery/models"
)

// ResultDao database result
type ResultDao struct {
	engine *xorm.Engine
}

// NewResultDao entrance result database
func NewResultDao(engine *xorm.Engine) *ResultDao {
	return &ResultDao{
		engine: engine,
	}
}

// Get Get by id
func (d *ResultDao) Get(id int) *models.LtResult {
	data := &models.LtResult{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	data.Id = 0
	return data
}

// GetAll Get some page datas
func (d *ResultDao) GetAll(page, size int) []models.LtResult {
	offset := (page - 1) * size
	datalist := make([]models.LtResult, 0)
	err := d.engine.
		Desc("id").
		Limit(size, offset).
		Find(&datalist)
	if err != nil {
		return datalist
	}
	return datalist
}

// CountAll CountAll for result
func (d *ResultDao) CountAll() int64 {
	num, err := d.engine.
		Count(&models.LtResult{})
	if err != nil {
		return 0
	}
	return num
}

// GetNewPrize TODO:
func (d *ResultDao) GetNewPrize(size int, giftIDs []int) []models.LtResult {
	datalist := make([]models.LtResult, 0)
	err := d.engine.
		In("gift_id", giftIDs).
		Desc("id").
		Limit(size).
		Find(&datalist)
	if err != nil {
		return datalist
	}
	return datalist
}

// SearchByGift ...
func (d *ResultDao) SearchByGift(giftID, page, size int) []models.LtResult {
	offset := (page - 1) * size
	datalist := make([]models.LtResult, 0)
	err := d.engine.
		Where("gift_id=?", giftID).
		Desc("id").
		Limit(size, offset).
		Find(&datalist)
	if err != nil {
		return datalist
	}
	return datalist
}

// SearchByUser ...
func (d *ResultDao) SearchByUser(uid, page, size int) []models.LtResult {
	offset := (page - 1) * size
	datalist := make([]models.LtResult, 0)
	err := d.engine.
		Where("uid=?", uid).
		Desc("id").
		Limit(size, offset).
		Find(&datalist)
	if err != nil {
		return datalist
	}
	return datalist

}

// CountByGift ...
func (d *ResultDao) CountByGift(giftID int) int64 {
	num, err := d.engine.
		Where("gift_id=?", giftID).
		Count(&models.LtResult{})
	if err != nil {
		return 0
	}
	return num
}

// CountByUser ...
func (d *ResultDao) CountByUser(uid int) int64 {
	num, err := d.engine.
		Where("uid=?", uid).
		Count(&models.LtResult{})
	if err != nil {
		return 0
	}
	return num

}
// Delete ...
func (d *ResultDao) Delete(id int) error {
	data := &models.LtResult{Id: id, SysStatus: 1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

// Update ...
func (d *ResultDao) Update(data *models.LtResult, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

// Create ...
func (d *ResultDao) Create(data *models.LtResult) error {
	_, err := d.engine.Insert(data)
	return err
}
