package models

type Userday struct {
	Day        int `xorm:"INT(11)"`
	Id         int `xorm:"not null pk autoincr INT(11)"`
	Num        int `xorm:"INT(11)"`
	SysCreated int `xorm:"INT(11)"`
	SysUpdated int `xorm:"INT(11)"`
	Uid        int `xorm:"INT(11)"`
}
