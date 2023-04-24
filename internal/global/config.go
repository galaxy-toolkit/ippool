package global

import "github.com/galaxy-toolkit/server/config"

// Config 配置
var Config config.Config

// InitConfig 初始化读取配置
func InitConfig(path string) {
	if err := config.LoadAndWatch(path, &Config); err != nil {
		panic(err)
	}
}
