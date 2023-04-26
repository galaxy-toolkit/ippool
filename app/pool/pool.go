package pool

import (
	"context"
	"github.com/galaxy-toolkit/ippool/domain/pool"
	"github.com/galaxy-toolkit/ippool/request/crawl/kuaidaili"
	"golang.org/x/exp/slog"
)

func Run(ctx context.Context) error {
	ips, err := kuaidaili.New(ctx).Crawl()
	if err != nil {
		slog.ErrorCtx(ctx, "Run Crawl err", err)
		return err
	}

	err = pool.Use(ctx).IP.InsertMany(ips)
	if err != nil {
		slog.ErrorCtx(ctx, "Run InsertMany err", err)
		return err
	}

	return nil
}
