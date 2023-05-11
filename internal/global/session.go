package global

import (
	"github.com/galaxy-toolkit/server/server"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// Session 存储
var Session *session.Store

// InitSessionStore 初始化 session 存储
func InitSessionStore() {
	Session = server.NewSessionStore(Config.Database.Redis)
}
