package bootstrap

import (
	"strconv"
	"time"
	"io/ioutil"

	"lottery/conf"
	"lottery/cron"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

type Configurator func(*Bootstrapper)

type Bootstrapper struct {
	*iris.Application
	AppName     string
	AppOwner    string
	AppSpawDate time.Time
}

func New(appName, appOwner string, cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		Application: iris.New(),
		AppName:     appName,
		AppOwner:    appOwner,
		AppSpawDate: time.Now(),
	}

	for _, cfg := range cfgs {
		cfg(b)
	}

	return b
}

// SetupViews reloads the templates.
func (b *Bootstrapper) SetupViews(viewsDir string) {
	// 全局模板视图
	htmlEngine := iris.HTML(viewsDir, ".html").Layout("shared/layout.html")
	// 每次重新加载模版（线上关闭它）
	htmlEngine.Reload(false)
	// 给模版内置各种定制的方法
	// 格式化时间
	htmlEngine.AddFunc("FromUnixtimeShort", func(t int) string {
		dt := time.Unix(int64(t), int64(0))
		return dt.Format(conf.SysTimeformShort)
	})
	htmlEngine.AddFunc("FromUnixtime", func(t int) string {
		dt := time.Unix(int64(t), int64(0))
		return dt.Format(conf.SysTimeform)
	})

	htmlEngine.AddFunc("PrePage", func(data int) string {
		pageIndex := strconv.Itoa(data - 1)
		return pageIndex
	})

	htmlEngine.AddFunc("NextPage", func(data int) string {
		pageIndex := strconv.Itoa(data + 1)
		return pageIndex
	})

	b.RegisterView(htmlEngine)
}

// SetupErrorHandlers response error handlers
func (b *Bootstrapper) SetupErrorHandlers() {
	b.OnAnyErrorCode(func(ctx iris.Context) {
		err := iris.Map{
			"app":     b.AppName,
			"status":  ctx.GetStatusCode(),
			"message": ctx.Values().GetString("message"),
		}

		if jsonOutput := ctx.URLParamExists("json"); jsonOutput {
			ctx.JSON(err)
			return
		}

		ctx.ViewData("Err", err)
		ctx.ViewData("Title", "Error")
		ctx.View("shared/error.html")
	})
}

// Configure accepts configurations and runs them inside the Bootstraper's context.
func (b *Bootstrapper) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}

func (b *Bootstrapper) setupCron() {
	cron.ConfigueAppOneCron()
}

const (
	// StaticAssets is the root directory for public assets like images, css, js.
	StaticAssets = "./public"
	// Favicon is the relative 9to the "StaticAssets") favicon path for our app.
	Favicon      = "/favicon.ico"
)

// Bootstrap init
func (b *Bootstrapper) Bootstrap() *Bootstrapper {
	b.SetupViews("./views")
	b.SetupErrorHandlers()

	// static files
	b.Favicon(StaticAssets + Favicon)
	b.StaticWeb(StaticAssets[1:], StaticAssets)
	indexHTML, err := ioutil.ReadFile(StaticAssets + "/index.html")
	if err == nil {
		b.StaticContent(StaticAssets[1:] + "/", "text/html",
			indexHTML)
	}
	// 不要把目录末尾"/"省略掉
	iris.WithoutPathCorrectionRedirection(b.Application)

	// middleware, after static files
	b.Use(recover.New())
	b.Use(logger.New())

	return b
}

// Listen starts the http server with the specified "addr".
func (b *Bootstrapper) Listen(addr string, cfgs ...iris.Configurator) {
	b.Run(iris.Addr(addr), cfgs...)
}
