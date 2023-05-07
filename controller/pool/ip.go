package pool

import (
	"github.com/galaxy-toolkit/ippool/domain/pool"
	"github.com/galaxy-toolkit/server/server"
	"github.com/galaxy-toolkit/server/server/code"
	"github.com/gofiber/fiber/v2"
)

// IPListRequest 请求参数
type IPListRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// IPList 获取 IP 列表
// @Summary 获取 IP 列表
// @Description 获取 IP 列表
// @Tags IP
// @Accept json
// @Produce json
// @Success 200 {object} server.DataResponse[server.PageResponse[[]*model.IP]]
// @Router /ip [get]
func IPList(ctx *fiber.Ctx) error {
	var req IPListRequest
	if err := ctx.BodyParser(&req); err != nil {
		return server.SendCode(ctx, code.ParamsParseFailed)
	}

	ips, total, err := pool.Use(ctx.Context()).IP.Find(req.Page, req.PageSize)
	if err != nil {
		return server.SendFailed(ctx)
	}

	return server.SendPageDataOk(ctx, ips, req.Page, req.PageSize, total)
}
