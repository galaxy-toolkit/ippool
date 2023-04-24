package main

import (
	"fmt"

	"github.com/galaxy-toolkit/ippool/internal/global"
)

func main() {
	global.InitConfig("config.yaml") // 配置
	global.InitLogger()              // 日志
	global.InitPostgres()            // Postgres

	fmt.Println(global.Config)

	global.Logger.Info("hello world")
}
