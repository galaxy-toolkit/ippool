package router

import (
	"github.com/galaxy-toolkit/ippool/internal/global"
	"github.com/galaxy-toolkit/server/server"
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
		server.NewLimiterHandler(global.Config.Server),
		server.NewLoggerHandler(global.Config.Server, global.LoggerWriter),
	)

	WithSwagger(app, global.Config.Server)
	WithPool(app)

	if err := app.Listen(global.Config.Server.Host + ":" + global.Config.Server.Port); err != nil {
		panic(err)
	}
}
