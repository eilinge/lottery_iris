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

// 用户信息，可以缓存(本地或者redis)，有更新的时候，可以直接清除缓存或者根据具体情况更新缓存
var cachedUserList = make(map[int]*models.LtUser)
var cachedUserLock = sync.Mutex{}

type UserService interface {
	GetAll(page, size int) []models.LtUser
	CountAll() int
	Get(id int) *models.LtUser
	Update(user *models.LtUser, columns []string) error
	Create(user *models.LtUser) error
	getByCache(id int) *models.LtUser
	setByCache(data *models.LtUser)
	updateByCache(data *models.LtUser, columns []string)
}

type userService struct {
	dao *dao.UserDao
}

func NewUserService() UserService {
	return &userService{
		dao: dao.NewUserDao(datasource.InstanceDbMaster()),
	}
}
// 只处理了单个数据, 不需要对列表进行缓存操作
func (s *userService) GetAll(page, size int) []models.LtUser {
	return s.dao.GetAll(page, size)
}

func (s *userService) CountAll() int {
	return s.dao.CountAll()
}

func (s *userService) Get(id int) *models.LtUser {
	data := s.getByCache(id)
	if data == nil || data.Id <= 0 {
		data := s.dao.Get(id)
		// 假使数据库中没有, 则将空数据放入缓存中, 方便从缓存中读
		if data == nil || data.Id <= 0 {
			data = &models.LtUser{Id: id}
		}
		s.setByCache(data)
	}
	
	return data
}

func (s *userService) Update(data *models.LtUser, columns []string) error {
	// 清空缓存, 然后再更新数据, 再进行读取时, 会获得新的数据, 更新到缓存
	s.updateByCache(data, columns)
	return s.dao.Update(data, columns)
}

func (s *userService) Create(data *models.LtUser) error {
    err := s.dao.Create(data)
  	return err
}

/*
单个用户数据的缓存, hash结构, user_service
hash: 直接对用户数据的某个字段进行更新, 不需要使用到json序列化

增加3个方法: getByCaChe, setByCaChe, updateByCache
修改Get, Update方法, 增加对缓存方法的调用
修改数据的时候, 清空缓存, 下次读取的时候, 自动更新最新数据
*/

// 从缓存中得到信息
func (s *userService) getByCache(id int) *models.LtUser {
	// 集群模式，redis缓存
	key := fmt.Sprintf("info_user_%d", id)
	rds := datasource.InstanceCache()
	dataMap, err := redis.StringMap(rds.Do("HGETALL", key))
	if err != nil {
		log.Println("user_service.getByCache HGETALL key=", key, ", error=", err)
		return nil
	}
	dataId := comm.GetInt64FromStringMap(dataMap, "Id", 0)
	if dataId <= 0 {
		return nil
	}
	data := &models.LtUser{
		Id:         int(dataId),
		Username:   comm.GetStringFromStringMap(dataMap, "Username", ""),
		Blacktime:  int(comm.GetInt64FromStringMap(dataMap, "Blacktime", 0)),
		Realname:   comm.GetStringFromStringMap(dataMap, "Realname", ""),
		Mobile:     comm.GetStringFromStringMap(dataMap, "Mobile", ""),
		Address:    comm.GetStringFromStringMap(dataMap, "Address", ""),
		SysCreated: int(comm.GetInt64FromStringMap(dataMap, "SysCreated", 0)),
		SysUpdated: int(comm.GetInt64FromStringMap(dataMap, "SysUpdated", 0)),
		SysIp:      comm.GetStringFromStringMap(dataMap, "SysIp", ""),
	}
	return data
}

func (s *userService) setByCache(data *models.LtUser) {
	if data.Id <= 0 || data == nil {
		return
	}
	id := data.Id
	key := fmt.Sprintf("info_user_%d", id)
	rds := datasource.InstanceCache()
	params := redis.Args{key}
	params = params.Add(id)
	if data.Username != "" {
		params = params.Add(params, "Username", data.Username)
		params = params.Add(params, "Blacktime", data.Blacktime)
		params = params.Add(params, "Realname", data.Realname)
		params = params.Add(params, "Mobile", data.Mobile)
		params = params.Add(params, "Address", data.Address)
		params = params.Add(params, "SysCreated", data.SysCreated)
		params = params.Add(params, "SysUpdated", data.SysUpdated)
		params = params.Add(params, "SysIp", data.SysIp)
	}
	_, err := rds.Do("HMSET", params)
	if err != nil {
		log.Println("user_service.setByCache HMSET params=", params, ", error=", err)
		return
	}
}

// 未对models.LtUser直接更新, 代码难度, 验证的精确性较高
func (s *userService) updateByCache(data *models.LtUser, columns []string) {
	if data == nil || data.Id <= 0 {
		return
	}
	key := fmt.Sprintf("info_user_%d", data.Id)
	rds := datasource.InstanceCache()
	_, err := rds.Do("DEL", key)
	if err != nil {
		log.Println("user_service.setByCache DEL key=", key, ", error=", err)
		return
	}
}