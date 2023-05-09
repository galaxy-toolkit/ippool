package pool

import (
	"github.com/galaxy-toolkit/ippool/domain/model"
	"github.com/galaxy-toolkit/ippool/domain/pool"
	"github.com/galaxy-toolkit/server/server"
	"github.com/galaxy-toolkit/server/server/code"
	"github.com/gofiber/fiber/v2"
)

// IPListRequest 请求参数
type IPListRequest struct {
	Page     int `json:"page" validate:"required,max=1000"`
	PageSize int `json:"page_size" validate:"required,max=20"`
}

// IPListResponse 响应
type IPListResponse server.DataResponse[server.PageResponse[[]*model.IP]]

// IPList 获取 IP 列表
//
//	@Summary		获取 IP 列表
//	@Description	获取 IP 列表
//	@Tags			IP
//	@Accept			json
//	@Produce		json
//	@Param			q	body		IPListRequest	true	"请求参数"
//	@Success		200	{object}	IPListResponse
//	@Failure		200	{object}	server.ParamsParseFailedResponse
//	@Router			/ip [get]
func IPList(ctx *fiber.Ctx) error {
	var req IPListRequest
	if err := ctx.BodyParser(&req); err != nil {
		return server.SendCode(ctx, code.ParamsParseFailed)
	}

	if err := server.Validate[IPListRequest](req); err != nil {
		return server.SendParamsParseFailed(ctx, err)
	}

	ips, total, err := pool.Use(ctx.Context()).IP.Find(req.Page, req.PageSize)
	if err != nil {
		return server.SendFailed(ctx)
	}

	return server.SendPageDataOk(ctx, ips, req.Page, req.PageSize, total)
}
