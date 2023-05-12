package router

import (
	"github.com/galaxy-toolkit/ippool/controller/user"
	"github.com/gofiber/fiber/v2"
)

func WithUser(router *fiber.App) {
	userGroup := router.Group("/user")
	{
		userGroup.Post("/register", user.Register)
		userGroup.Post("/login", user.Login)
		userGroup.Get("/:id", user.GetUserInfo)
	}
}
