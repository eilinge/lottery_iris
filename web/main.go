package main

import (
	"imooc.com/lottery/bootstrap"
	"imooc.com/lottery/web/middleware/identity"
	"imooc.com/lottery/web/routes"
	"fmt"
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
	app := newApp()
	app.Listen(fmt.Sprintf(":%d", port))
}
