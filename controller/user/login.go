package user

import (
	"github.com/galaxy-toolkit/ippool/domain/model"
	"github.com/galaxy-toolkit/ippool/domain/user"
	"github.com/galaxy-toolkit/ippool/internal/global"
	"github.com/galaxy-toolkit/server/server"
	"github.com/galaxy-toolkit/server/server/code"
	"github.com/gofiber/fiber/v2"
)

// LoginRequestParams 登录请求参数
type LoginRequestParams struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginResponse 响应
type LoginResponse server.DataResponse[*model.User]

// Login 登录
//
//	@Summary		登录
//	@Description	登录
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			q	body		LoginRequestParams	true	"请求参数"
//	@Success		200	{object}	LoginResponse
//	@Router			/user/login [post]
func Login(ctx *fiber.Ctx) error {
	var req LoginRequestParams
	if err := ctx.BodyParser(&req); err != nil {
		return server.SendCode(ctx, code.ParamsParseFailed)
	}

	if err := server.Validate[LoginRequestParams](req); err != nil {
		return server.SendParamsParseFailed(ctx, err)
	}

	u, err := user.Use(ctx.Context()).User.FindOne(&user.FindOneOption{Name: req.Username})
	if err != nil {
		return server.SendCode(ctx, code.UserNotFound)
	}
	if u == nil {
		return server.SendCode(ctx, code.UserNotFound)
	}

	if req.Password != u.Password {
		return server.SendCode(ctx, code.PasswordError)
	}

	sess, err := global.Session.Get(ctx)
	if err != nil {
		return server.SendFailed(ctx)
	}

	sess.Set("username", u.Username)
	if err := sess.Save(); err != nil {
		return server.SendFailed(ctx)
	}

	return server.SendDataOk(ctx, u)
}
