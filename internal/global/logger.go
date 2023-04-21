package global

import (
	"github.com/galaxy-toolkit/server/log"
	"golang.org/x/exp/slog"
)

// Logger 日志
var Logger *slog.Logger

// InitLogger 初始化日志
func InitLogger() {
	logger, err := log.New(Config.Log)
	if err != nil {
		panic(err)
	}
	Logger = logger
}
