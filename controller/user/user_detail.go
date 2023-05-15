package user

import (
	"strconv"

	"github.com/galaxy-toolkit/ippool/app/user"
	"github.com/galaxy-toolkit/ippool/domain/model"
	"github.com/galaxy-toolkit/server/log"
	"github.com/galaxy-toolkit/server/server"
	"github.com/galaxy-toolkit/server/server/code"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
)

// GetUserInfoResponse 响应
type GetUserInfoResponse server.DataResponse[*model.User]

// GetUserInfo 用户详情
//
//	@Summary		用户详情
//	@Description	用户详情
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"用户 ID"
//	@Success		200	{object}	GetUserInfoResponse
//	@Router			/user/{id} [get]
func GetUserInfo(ctx *fiber.Ctx) error {
	uidParams := ctx.Params("id")
	if uidParams == "" {
		log.Server.Error(ctx.Context(), "参数解析失败")
		return server.SendCode(ctx, code.ParamsParseFailed)
	}
	uid, err := strconv.ParseInt(uidParams, 10, 64)
	if err != nil {
		log.Server.Error(ctx.Context(), "参数验证失败", slog.Any("err", err))
		return server.SendCode(ctx, code.ParamsParseFailed)
	}

	u, err := user.NewService(ctx.Context()).Detail(uid)
	if err != nil {
		log.Server.Error(ctx.Context(), "查询用户详情失败", slog.Any("err", err))
		return server.SendError(ctx, err)
	}

	return server.SendDataOk(ctx, u)
}
