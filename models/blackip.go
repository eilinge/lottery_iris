package models

type Blackip struct {
	Address    string `xorm:"VARCHAR(255)"`
	Blacktime  int    `xorm:"INT(11)"`
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	Mobile     string `xorm:"VARCHAR(50)"`
	Realname   string `xorm:"VARCHAR(50)"`
	SysCreated int    `xorm:"INT(11)"`
	SysIp      string `xorm:"VARCHAR(50)"`
	SysUpdated int    `xorm:"INT(11)"`
	Username   string `xorm:"VARCHAR(255)"`
}
