package router

import (
	swag "github.com/galaxy-toolkit/ippool/internal/swagger"
	"github.com/galaxy-toolkit/server/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// WithSwagger 配置 swagger 文档接口
func WithSwagger(app *fiber.App, conf config.Server) {
	if !conf.EnableSwagger {
		return
	}

	swag.SwaggerInfo.Title = conf.Title
	swag.SwaggerInfo.Host = conf.Host + ":" + conf.Port
	swag.SwaggerInfo.BasePath = conf.BasePath
	swag.SwaggerInfo.Version = conf.Version
	app.Get("/swagger/*", swagger.New(swagger.Config{}))
}
