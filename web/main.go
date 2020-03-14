package main

import (
	"fmt"
	"imooc.com/lottery/bootstrap"
	"imooc.com/lottery/conf"
	"imooc.com/lottery/web/middleware/identity"
	"imooc.com/lottery/web/routes"
)

var port = 8000

func newApp() *bootstrap.Bootstrapper {
	// 初始化应用
	app := bootstrap.New("Go抽奖系统", "一凡Sir")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)

	return app
}

func main() {
	// 服务器集群的时候才需要区分这项设置
	// 比如：根据服务器的IP、名称、端口号等，或者运行的参数
	if port == 8000 {
		conf.RunningCrontabService = true
	}

	app := newApp()
	app.Listen(fmt.Sprintf(":%d", port))
}
