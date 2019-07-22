package dao

/*
 抽奖系统的数据库操作
*/

import (
	"github.com/go-xorm/xorm"

	"lottery/models"
)

// CodeDao Database code
type CodeDao struct {
	engine *xorm.Engine
}

// NewCodeDao get new *CodeDao
func NewCodeDao(engine *xorm.Engine) *CodeDao {
	return &CodeDao{
		engine: engine,
	}
}

// Get get code by id
func (d *CodeDao) Get(id int) *models.LtCode {
	data := &models.LtCode{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	data.Id = 0
	return data
}

// GetAll get all from code
func (d *CodeDao) GetAll(page, size int) []models.LtCode {
	offset := (page - 1) * size
	datalist := make([]models.LtCode, 0)
	err := d.engine.
		Desc("id").
		Limit(size, offset).
		Find(&datalist)
	if err != nil {
		return datalist
	}
	return datalist
}

// CountAll count database code
func (d *CodeDao) CountAll() int64 {
	num, err := d.engine.
		Count(&models.LtCode{})
	if err != nil {
		return 0
	}
	return num
}

// CountByGift ...
func (d *CodeDao) CountByGift(giftID int) int64 {
	num, err := d.engine.
		Where("gift_id=?", giftID).
		Count(&models.LtCode{})
	if err != nil {
		return 0
	}
	return num

}

// Search by giftID
func (d *CodeDao) Search(giftID int) []models.LtCode {
	datalist := make([]models.LtCode, 0)
	err := d.engine.
		Where("gift_id=?", giftID).
		Desc("id").
		Find(&datalist)
	if err != nil {
		return datalist
	}
	return datalist
}

// Delete delete by id
func (d *CodeDao) Delete(id int) error {
	data := &models.LtCode{Id: id, SysStatus: 1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

// Update update more columns
func (d *CodeDao) Update(data *models.LtCode, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

// Create insert new code
func (d *CodeDao) Create(data *models.LtCode) error {
	_, err := d.engine.Insert(data)
	return err
}

// NextUsingCode 找到下一个可用的最小的优惠券
func (d *CodeDao) NextUsingCode(giftID, codeID int) *models.LtCode {
	datalist := make([]models.LtCode, 0)
	err := d.engine.Where("gift_id=?", giftID).
		Where("sys_status=?", 0).
		Where("id>?", codeID).
		Asc("id").Limit(1).
		Find(&datalist)
	if err != nil || len(datalist) < 1 {
		return nil
	}
	return &datalist[0]

}

// UpdateByCode 根据唯一的code来更新
func (d *CodeDao) UpdateByCode(data *models.LtCode, columns []string) error {
	_, err := d.engine.Where("code=?", data.Code).
		MustCols(columns...).Update(data)
	return err
}
