package main

import (
	"github.com/galaxy-toolkit/ippool/internal/global"
	"github.com/galaxy-toolkit/ippool/router"
)

func main() {
	global.InitConfig("config.yaml") // 配置
	global.InitLogger()              // 日志
	global.InitPostgres()            // Postgres

	router.Run()

	global.Logger.Info("end")
}
