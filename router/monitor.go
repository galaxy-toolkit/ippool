package router

import (
	"github.com/galaxy-toolkit/ippool/internal/global"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func WithMonitor(router *fiber.App) {
	if global.Config.Server.Monitor {
		router.Get("/metrics", monitor.New(monitor.Config{
			Title: global.Config.Server.Title,
		}))
	}
}
