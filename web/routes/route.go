package routes

import (
	"github.com/kataras/iris/mvc"

	"lottery/bootstrap"
	"lottery/services"
	"lottery/web/controllers"
	"lottery/web/middleware"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {
	userService := services.NewUserService()
	giftService := services.NewGiftService()
	codeService := services.NewCodeService()
	resultService := services.NewResultService()
	userdayService := services.NewUserdayService()
	blackipService := services.NewBlackipService()

	// 路由分组IndexController/AdminController
	// "/" IndexController
	index := mvc.New(b.Party("/"))
	index.Register(userService,
		giftService,
		codeService,
		resultService,
		userdayService,
		blackipService)
	index.Handle(new(controllers.IndexController))

	// "/admin" AdminController
	admin := mvc.New(b.Party("/admin"))
	admin.Router.Use(middleware.BasicAuth)
	admin.Register(userService,
		giftService,
		codeService,
		resultService,
		userdayService,
		blackipService)
	admin.Handle(new(controllers.AdminController))

	// "/admin/gift" AdminController
	adminGift := admin.Party("/gift")
	adminGift.Register(giftService)
	adminGift.Handle(new(controllers.AdminGiftController))

	// 传统设置路由
	//b.Get("/follower/{id:long}", GetFollowerHandler)
	//b.Get("/following/{id:long}", GetFollowingHandler)
	//b.Get("/like/{id:long}", GetLikeHandler)
}
