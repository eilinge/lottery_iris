package models

type Result struct {
	GiftData   string `xorm:"VARCHAR(255)"`
	GiftId     int    `xorm:"INT(11)"`
	GiftName   string `xorm:"VARCHAR(255)"`
	GiftType   int    `xorm:"INT(11)"`
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	PrizeCode  int    `xorm:"INT(11)"`
	SysCreated int    `xorm:"INT(11)"`
	SysIp      int    `xorm:"INT(11)"`
	SysStatus  int    `xorm:"SMALLINT(6)"`
	Uid        int    `xorm:"INT(11)"`
	Username   string `xorm:"VARCHAR(255)"`
}
