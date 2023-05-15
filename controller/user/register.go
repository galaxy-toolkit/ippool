package user

import (
	"github.com/galaxy-toolkit/ippool/app/user"
	"github.com/galaxy-toolkit/server/log"
	"github.com/galaxy-toolkit/server/server"
	"github.com/galaxy-toolkit/server/server/code"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
)

// RegisterRequestParams 注册请求参数
type RegisterRequestParams struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Register 注册
//
//	@Summary		注册
//	@Description	注册
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			q	body		RegisterRequestParams	true	"请求参数"
//	@Success		200	{object}	server.BasicResponse
//	@Router			/user/register [post]
func Register(ctx *fiber.Ctx) error {
	var req RegisterRequestParams
	if err := ctx.BodyParser(&req); err != nil {
		log.Server.Error(ctx.Context(), "参数解析失败", slog.Any("err", err))
		return server.SendCode(ctx, code.ParamsParseFailed)
	}

	if err := server.Validate[RegisterRequestParams](req); err != nil {
		log.Server.Error(ctx.Context(), "参数验证失败", slog.Any("err", err))
		return server.SendParamsParseFailed(ctx, err)
	}

	err := user.NewService(ctx.Context()).Register(&user.RegisterParams{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		log.Server.Error(ctx.Context(), "注册失败", slog.Any("err", err))
		return server.SendError(ctx, err)
	}

	return server.SendOk(ctx)
}
