package main

import (
	"imooc.com/lottery/bootstrap"
	"imooc.com/lottery/web/middleware/identity"
	"imooc.com/lottery/web/routes"
	"fmt"
	"imooc.com/lottery/conf"
	)

var port = 8080

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
	if port == 8080 {
		conf.RunningCrontabService = true
	}

	// TODO: 定义thrift方案，以及处理前端大转盘页面

	app := newApp()
	app.Listen(fmt.Sprintf(":%d", port))
}
