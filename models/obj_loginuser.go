package models

// ObjLoginuser 站点中与浏览器交互的用户模型
type ObjLoginuser struct {
	Uid      int	`json:"uid"`
	Username string	`json:"username"`
	Now      int	`json:"now"`
	Ip       string	`json:"ip"`
	Sign     string	`json:"sign"`
}
