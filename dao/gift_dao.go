package dao

/*
 抽奖系统的数据库操作
*/

import (
	"log"
	"github.com/go-xorm/xorm"

	"lottery/models"
	"lottery/comm"
)

// GiftDao Database gift
type GiftDao struct {
	engine *xorm.Engine
}

// NewGiftDao get new *GiftDao
func NewGiftDao(engine *xorm.Engine) *GiftDao {
	return &GiftDao{
		engine: engine,
	}
}

// Get get gift by id
func (d *GiftDao) Get(id int) *models.LtGift {
	data := &models.LtGift{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	log.Println("failed to get err ", err)
	data.Id = 0
	return data

}

// GetAll get all from gift
func (d *GiftDao) GetAll() []models.LtGift {
	datalist := make([]models.LtGift, 0)
	err := d.engine.
		Asc("sys_status").
		Asc("displayorder").
		Find(&datalist)
	if err != nil {
		log.Println("failed to GetAll ", err)
		return datalist
	}
	return datalist

}

// CountAll count database gift
func (d *GiftDao) CountAll() int64 {
	num, err := d.engine.
		Count(&models.LtGift{})
	if err != nil {
		return 0
	}
	return num

}

// Delete delete gift by id
func (d *GiftDao) Delete(id int) error {
	data := &models.LtGift{Id: id, SysStatus: 1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

// Update update more columns
func (d *GiftDao) Update(data *models.LtGift, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

// Create insert new gift
func (d *GiftDao) Create(data *models.LtGift) error {
	_, err := d.engine.Insert(data)
	return err
}

// GetAllUse ...
// 获取到当前可以获取的奖品列表
// 有奖品限定，状态正常，时间期间内
// gtype倒序， displayorder正序
func (d *GiftDao) GetAllUse() []models.LtGift {
	now := comm.NowUnix()
	datalist := make([]models.LtGift, 0)
	err := d.engine.
		Cols("id", "title", "prize_num", "left_num", "prize_code",
			"prize_time", "img", "displayorder", "gtype", "gdata").
		Desc("gtype").
		Asc("displayorder").
		Where("prize_num>=?", 0). 		// 有限定的奖品
		Where("sys_status=?", 0). 		// 有效的奖品
		Where("time_begin<=?", now).   	// 时间期内
		Where("time_end>=?", now).     	// 时间期内
		Find(&datalist)
	if err != nil {
		return datalist
	}
	return datalist
}

// IncrLeftNum increase gift left number
func (d *GiftDao) IncrLeftNum(id, num int) (int64, error) {
	r, err := d.engine.Id(id).
		Incr("left_num", num).
		Update(&models.LtGift{Id:id})
	return r, err
}

// DecrLeftNum decrease gift left number
func (d *GiftDao) DecrLeftNum(id, num int) (int64, error) {
	r, err := d.engine.Id(id).
		Decr("left_num", num).
		Where("left_num>=?", num).
		Update(&models.LtGift{Id:id})
	return r, err
}
