package dao

import (
	"github.com/go-xorm/xorm"

	"lottery/models"
)

// UserdayDao database userday
type UserdayDao struct {
	engine *xorm.Engine
}

// NewUserdayDao entrance userdaydao
func NewUserdayDao(engine *xorm.Engine) *UserdayDao {
	return &UserdayDao{
		engine: engine,
	}
}

// Get Get by id
func (d *UserdayDao) Get(id int) *models.LtUserday {
	data := &models.LtUserday{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	data.Id = 0
	return data
}

// GetAll Get some page datas
func (d *UserdayDao) GetAll(page, size int) []models.LtUserday {
	offset := (page - 1) * size
	datalist := make([]models.LtUserday, 0)
	err := d.engine.
		Desc("id").
		Limit(size, offset).
		Find(&datalist)
	if err != nil {
		return datalist
	}
	return datalist
}

// CountAll count all
func (d *UserdayDao) CountAll() int64 {
	num, err := d.engine.
		Count(&models.LtUserday{})
	if err != nil {
		return 0
	}
	return num
}

// Search Search userday by uid on today
func (d *UserdayDao) Search(uid, day int) []models.LtUserday {
	datalist := make([]models.LtUserday, 0)
	err := d.engine.
		Where("uid=?", uid).
		Where("day=?", day).
		Desc("id").
		Find(&datalist)
	if err != nil {
		return nil
	}
	return datalist
}

// Count get userday num by uid on today
func (d *UserdayDao) Count(uid, day int) int {
	info := &models.LtUserday{}
	ok, err := d.engine.
		Where("uid=?", uid).
		Where("day=?", day).
		Get(info)
	if !ok || err != nil {
		return 0
	}
	return info.Num

}

// Update Update userday
func (d *UserdayDao) Update(data *models.LtUserday, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

// Create Create a new userday
func (d *UserdayDao) Create(data *models.LtUserday) error {
	_, err := d.engine.Insert(data)
	return err
}
