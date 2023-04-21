package global

import (
	"io"

	"github.com/galaxy-toolkit/server/log"
	"golang.org/x/exp/slog"
)

// Logger 日志
var Logger *slog.Logger
var LoggerWriter io.Writer

// InitLogger 初始化日志
func InitLogger() {
	LoggerWriter = log.Writer(Config.Log)

	logger, err := log.New(Config.Log, LoggerWriter)
	if err != nil {
		panic(err)
	}
	Logger = logger
}
