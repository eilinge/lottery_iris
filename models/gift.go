package models

type Gift struct {
	Displayorder int    `xorm:"INT(11)"`
	Gdata        string `xorm:"VARCHAR(255)"`
	Gtype        int    `xorm:"INT(11)"`
	Id           int    `xorm:"not null pk autoincr INT(11)"`
	Img          string `xorm:"VARCHAR(255)"`
	LeftNum      int    `xorm:"INT(11)"`
	PrizeBegin   int    `xorm:"INT(11)"`
	PrizeCode    string `xorm:"VARCHAR(50)"`
	PrizeData    string `xorm:"MEDIUMTEXT"`
	PrizeEnd     int    `xorm:"INT(11)"`
	PrizeNum     int    `xorm:"INT(11)"`
	PrizeTime    int    `xorm:"INT(11)"`
	SysCreated   int    `xorm:"INT(11)"`
	SysIp        string `xorm:"VARCHAR(50)"`
	SysStatus    int    `xorm:"SMALLINT(6)"`
	SysUpdated   int    `xorm:"INT(11)"`
	TimeBegin    int    `xorm:"INT(11)"`
	TimeEnd      int    `xorm:"INT(11)"`
	Title        string `xorm:"VARCHAR(255)"`
}
