package main

import (
	"github.com/galaxy-toolkit/ippool/internal/global"
	"gorm.io/gen"
)

func main() {
	global.InitConfig() // 配置
	global.InitLogger() // 日志

	genIP()
}

func genIP() {
	generator, err := global.InitMySQLGenerator("./domain/query/ip", "", "./domain/model/ip")
	if err != nil {
		panic(err)
	}

	generator.GenerateModelAs("ip", "IP", gen.FieldType("status", "Status"))
	generator.Execute()
}
