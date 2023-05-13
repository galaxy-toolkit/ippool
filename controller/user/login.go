package user

import (
	"github.com/galaxy-toolkit/ippool/app/user"
	"github.com/galaxy-toolkit/ippool/domain/model"
	"github.com/galaxy-toolkit/ippool/internal/global"
	"github.com/galaxy-toolkit/server/log"
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
		log.Error(ctx.Context(), "登录参数解析失败", "err", err)
		return server.SendCode(ctx, code.ParamsParseFailed)
	}

	if err := server.Validate[LoginRequestParams](req); err != nil {
		log.Error(ctx.Context(), "登录参数验证失败", "err", err)
		return server.SendParamsParseFailed(ctx, err)
	}

	u, err := user.NewService(ctx.Context()).Login(&user.LoginParams{
		Username: req.Username,
		Password: req.Password,
		Phone:    "",
	})
	if err != nil {
		log.Error(ctx.Context(), "登陆失败", "err", err)
		return server.SendError(ctx, err)
	}

	if err := SetSession(ctx, u); err != nil {
		log.Error(ctx.Context(), "生成 session 失败", "err", err)
		return server.SendError(ctx, err)
	}

	return server.SendDataOk(ctx, u)
}

func SetSession(ctx *fiber.Ctx, u *model.User) error {
	sess, err := global.Session.Get(ctx)
	if err != nil {
		return err
	}

	sess.Set("username", u.Username)

	if err := sess.Save(); err != nil {
		return err
	}
	return nil
}
