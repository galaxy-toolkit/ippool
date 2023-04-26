package main

import (
	"context"
	"github.com/galaxy-toolkit/ippool/app/pool"
	"github.com/galaxy-toolkit/ippool/internal/global"
)

func main() {
	global.InitConfig("config.yaml") // 配置
	global.InitLogger()              // 日志
	global.InitPostgres()            // Postgres

	pool.Run(context.TODO())

	global.Logger.Info("hello world")
}
