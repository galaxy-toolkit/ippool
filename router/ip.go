package router

import (
	"github.com/galaxy-toolkit/ippool/controller/pool"
	"github.com/gofiber/fiber/v2"
)

// WithPool 配置 IP 池路由
func WithPool(app *fiber.App) {
	ipRouter := app.Group("/ip")
	{
		ipRouter.Get("", pool.IPList)
		ipRouter.Get("/random", pool.IPRandom)
	}
}
