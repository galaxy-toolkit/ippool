package main

import (
	"github.com/galaxy-toolkit/ippool/internal/global"
	"github.com/galaxy-toolkit/ippool/router"
	"github.com/galaxy-toolkit/server/log"
)

func main() {
	// basic init
	global.InitConfig("config.yaml")  // 配置
	log.InitLogger(global.Config.Log) // 日志

	// database init
	global.InitPostgres()     // Postgres
	global.InitRedis()        // Redis
	global.InitSessionStore() // Session Store

	router.Run()
}
