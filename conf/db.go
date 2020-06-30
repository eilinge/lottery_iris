/**
 * mysql数据库配置信息
 */
package conf

const DriverName = "mysql"

type DbConfig struct {
	Host      string
	Port      int
	User      string
	Pwd       string
	Database  string
	IsRunning bool // 是否正常运行
}
