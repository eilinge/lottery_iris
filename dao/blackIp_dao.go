package dao

import (
	"github.com/go-xorm/xorm"

	"lottery/models"
)

// BlackipDao ...
type BlackipDao struct {
	engine *xorm.Engine
}

// NewBlackipDao get new  BlackipDao
func NewBlackipDao(engine *xorm.Engine) *BlackipDao {
	return &BlackipDao{
		engine: engine,
	}
}

// Get get blackip by id
func (d *BlackipDao) Get(id int) *models.LtBlackip {
	data := &models.LtBlackip{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	data.Id = 0
	return data
}

// GetAll get all from blackip
func (d *BlackipDao) GetAll(page, size int) []models.LtBlackip {
	offset := (page - 1) * size
	datalist := make([]models.LtBlackip, 0)
	err := d.engine.
		Desc("id").
		Limit(size, offset).
		Find(&datalist)
	if err != nil {
		return datalist
	}
	return datalist
}

// CountAll count database Blackip
func (d *BlackipDao) CountAll() int64 {
	num, err := d.engine.
		Count(&models.LtBlackip{})
	if err != nil {
		return 0
	}
	return num
}

// Search by ip
func (d *BlackipDao) Search(ip string) []models.LtBlackip {
	datalist := make([]models.LtBlackip, 0)
	err := d.engine.
		Where("ip=?", ip).
		Desc("id").
		Find(&datalist)
	if err != nil {
		return datalist
	}
	return datalist
}

// Update update more columns
func (d *BlackipDao) Update(data *models.LtBlackip, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

// Create insert new code
func (d *BlackipDao) Create(data *models.LtBlackip) error {
	_, err := d.engine.Insert(data)
	return err
}

// GetByIP 根据IP获取信息
func (d *BlackipDao) GetByIP(ip string) *models.LtBlackip {
	datalist := make([]models.LtBlackip, 0)
	err := d.engine.
		Where("ip=?", ip).
		Desc("id").
		Limit(1).
		Find(&datalist)
	if err != nil || len(datalist) < 1 {
		return nil
	}
	return &datalist[0]
}
