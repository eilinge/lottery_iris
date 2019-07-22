package datasource

import (
	"fmt"
	"log"
	"lottery/conf"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var dbLock sync.Mutex
var masterInstance *xorm.Engine

func InstanceDbMaster() *xorm.Engine {
	if masterInstance != nil {
		return masterInstance
	}
	dbLock.Lock()
	defer dbLock.Unlock()
	// 防止后边还在排队的人进来之后, 再次创建, 确保masterInstance只被创建一次
	if masterInstance != nil {
		return masterInstance
	}

	return NewDbMaster()

}

func NewDbMaster() *xorm.Engine {
	sourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		conf.DbMaster.User,
		conf.DbMaster.Pwd,
		conf.DbMaster.Host,
		conf.DbMaster.Port,
		conf.DbMaster.Database)

	instance, err := xorm.NewEngine(conf.DriverName, sourceName)
	if err != nil {
		log.Fatal("dbhelper.NewDbMaster NewEngine error ", err)
		return nil
	}
	instance.ShowSQL(true)
	masterInstance = instance
	return instance
}
