/**
 * 抽奖系统数据处理（包括数据库，也包括缓存等其他形式数据）
 */
package services

import (
	"fmt"
	"log"
	"sync"

	"lottery/comm"
	"lottery/dao"
	"lottery/datasource"
	"lottery/models"
	
	"github.com/gomodule/redigo/redis"
)

// IP信息，可以缓存(本地或者redis)，有更新的时候，再根据具体情况更新缓存
var cachedBlackipList = make(map[string]*models.LtBlackip)
var cachedBlackipLock = sync.Mutex{}

// BlackipService interface methods
type BlackipService interface {
	GetAll(page, size int) []models.LtBlackip
	CountAll() int64
	Search(ip string) []models.LtBlackip
	Get(id int) *models.LtBlackip
	Update(blackip *models.LtBlackip, columns []string) error
	Create(blackip *models.LtBlackip) error
	GetByIp(ip string) *models.LtBlackip
}

type blackipService struct {
	dao *dao.BlackipDao
}

// NewBlackipService BlackipService entance 
func NewBlackipService() BlackipService {
	return &blackipService{
		dao: dao.NewBlackipDao(datasource.InstanceDbMaster()),
	}
}

func (s *blackipService) GetAll(page, size int) []models.LtBlackip {
	return s.dao.GetAll(page, size)
}

func (s *blackipService) CountAll() int64 {
	return s.dao.CountAll()
}

func (s *blackipService) Search(ip string) []models.LtBlackip {
	return s.dao.Search(ip)
}

func (s *blackipService) Get(id int) *models.LtBlackip {
	return s.dao.Get(id)
}

func (s *blackipService) Update(data *models.LtBlackip, columns []string) error {
	s.updateByCache(data, columns)
	// 再更新数据的数据
	return s.dao.Update(data, columns)
}

func (s *blackipService) Create(data *models.LtBlackip) error {
	return s.dao.Create(data)
}

// 根据IP读取IP的黑名单数据
func (s *blackipService) GetByIp(ip string) *models.LtBlackip {
	data := s.getByCache(ip)
	if data == nil || data.Ip == "" {
		data = s.dao.GetByIP(ip)
		if data == nil || data.Ip == "" {
			// 设置空数据
			data = &models.LtBlackip{Ip: ip}
		}
		s.setByCache(data)
	}
	
	return data
}

/*
Ip黑名单数据的缓存, hash结构, blackip_service

增加了3个方法: getByCache, setByCache, updateByCache
修改GetByIp, Update方法, 增加对缓存方法的调用
修改数据的时候, 清空缓存, 下次读取的时候自动更新最新数据
*/

// 从缓存中得到信息
func (s *blackipService) getByCache(ip string) *models.LtBlackip {
	// 集群模式，redis缓存
	key := fmt.Sprintf("info_blackip_%s", ip)
	rds := datasource.InstanceCache()
	dataMap, err := redis.StringMap(rds.Do("HGETALL", key))
	if err != nil {
		log.Println("blackip_service.getByCache HGETALL key=", key, ", error=", err)
		return nil
	}
	dataIP := comm.GetStringFromStringMap(dataMap, "Ip", "")
	if dataIP == "" {
		return nil
	}

	data := &models.LtBlackip{
		Id:         int(comm.GetInt64FromStringMap(dataMap, "Id", 0)),
		Ip:   dataIP,
		Blacktime:  int(comm.GetInt64FromStringMap(dataMap, "Blacktime", 0)),
		SysCreated: int(comm.GetInt64FromStringMap(dataMap, "SysCreated", 0)),
		SysUpdated: int(comm.GetInt64FromStringMap(dataMap, "SysUpdated", 0)),
	}
	return data
}

func (s *blackipService) setByCache(data *models.LtBlackip) {
	if data.Ip == "" || data == nil {
		return
	}
	ip := data.Ip
	key := fmt.Sprintf("info_blackip_%s", ip)
	rds := datasource.InstanceCache()
	params := []interface{}{key}
	params = append(params, "Ip", data.Ip)
	if data.Id > 0 {
		params = append(params, "Blacktime", data.Blacktime)
		params = append(params, "SysCreated", data.SysCreated)
		params = append(params, "SysUpdated", data.SysUpdated)
	}
	_, err := rds.Do("HMSET", params...)
	if err != nil {
		log.Println("blackip_service.setByCache HMSET params=", params, ", error=", err)
		return
	}
}

// 未对models.LtBlackip直接更新, 代码难度, 验证的精确性较高
func (s *blackipService) updateByCache(data *models.LtBlackip, columns []string) {
	if data == nil || data.Ip == "" {
		return
	}
	key := fmt.Sprintf("info_blackip_%s", data.Ip)
	rds := datasource.InstanceCache()
	_, err := rds.Do("DEL", key)
	if err != nil {
		log.Println("blackip_service.updateByCache DEL key=", key, ", error=", err)
		return
	}
}