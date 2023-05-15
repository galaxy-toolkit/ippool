package main

import (
	"github.com/galaxy-toolkit/ippool/internal/global"
	"github.com/galaxy-toolkit/server/database/postgres"
	"github.com/galaxy-toolkit/server/log"
	"gorm.io/gen"
)

func main() {
	global.InitConfig("config.yaml")  // 配置
	log.InitLogger(global.Config.Log) // 日志

	genIP()
}

func genIP() {
	generator, err := global.InitPostgresGenerator(postgres.GeneratorConfig{
		ModelPath: "domain/model",
	})
	if err != nil {
		panic(err)
	}

	generator.GenerateModelAs("ip", "IP",
		gen.FieldType("status", "IPStatus"),
		gen.FieldType("protocol", "IPProtocol"),
	)
	generator.GenerateModelAs("user", "User")
	generator.Execute()
}
