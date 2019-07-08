package models

type User struct {
	Blacktime  int    `xorm:"INT(11)"`
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	Ip         string `xorm:"VARCHAR(50)"`
	SysCreated int    `xorm:"INT(11)"`
	SysUpdated int    `xorm:"INT(11)"`
}
