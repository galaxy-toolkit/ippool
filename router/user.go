package router

import (
	"github.com/galaxy-toolkit/ippool/controller/user"
	"github.com/gofiber/fiber/v2"
)

func WithUser(router *fiber.App) {
	userGroup := router.Group("/user")
	{
		userGroup.Post("/login", user.Login)
	}
}
