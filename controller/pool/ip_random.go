package pool

import (
	"github.com/galaxy-toolkit/ippool/domain/model"
	"github.com/galaxy-toolkit/ippool/domain/pool"
	"github.com/galaxy-toolkit/server/server"
	"github.com/galaxy-toolkit/server/server/code"
	"github.com/gofiber/fiber/v2"
)

// IPRandomRequest 请求参数
type IPRandomRequest struct {
	PageSize int `json:"page_size" validate:"required,max=20"`
}

// IPRandomResponse 响应
type IPRandomResponse server.DataResponse[[]*model.IP]

// IPRandom 随机获取一批 IP
//
//	@Summary		随机获取一批 IP
//	@Description	随机获取一批 IP
//	@Tags			IP
//	@Accept			json
//	@Produce		json
//	@Param			q	body		IPRandomRequest	true	"请求参数"
//	@Success		200	{object}	IPRandomResponse
//	@Router			/ip/random [get]
func IPRandom(ctx *fiber.Ctx) error {
	var req IPRandomRequest
	if err := ctx.BodyParser(&req); err != nil {
		return server.SendCode(ctx, code.ParamsParseFailed)
	}

	if err := server.Validate[IPRandomRequest](req); err != nil {
		return server.SendParamsParseFailed(ctx, err)
	}

	ips, err := pool.Use(ctx.Context()).IP.Random(req.PageSize)
	if err != nil {
		return server.SendFailed(ctx)
	}

	return server.SendDataOk(ctx, ips)
}
