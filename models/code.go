package models

type Code struct {
	Code       string `xorm:"VARCHAR(255)"`
	GiftId     int    `xorm:"INT(11)"`
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	SysCreated int    `xorm:"INT(11)"`
	SysStatus  int    `xorm:"SMALLINT(6)"`
	SysUpdated int    `xorm:"INT(11)"`
}
