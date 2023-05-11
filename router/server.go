package router

import (
	"github.com/galaxy-toolkit/ippool/internal/global"
	"github.com/galaxy-toolkit/server/server"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

// Run 启动服务
//
//	@title			ippool
//	@version		1.0
//	@description	ippool
//	@host			localhost:9999
//	@BasePath		/
func Run() {
	app := server.NewRouter()
	app.Use(
		recover.New(recover.Config{EnableStackTrace: true}), // Panic recover
		requestid.New(), // RequestID
		server.NewLoggerHandler(global.Config.Server, global.LoggerWriter), // 日志
		server.NewLimiterHandler(global.Config.Server),                     // 限流器
	)

	WithPool(app) // IP 池
	WithUser(app) // 用户

	if err := app.Listen(global.Config.Server.Host + ":" + global.Config.Server.Port); err != nil {
		panic(err)
	}
}
