package pool

import (
	"context"

	"github.com/galaxy-toolkit/ippool/domain/pool"
	"github.com/galaxy-toolkit/ippool/internal/global"
	"github.com/galaxy-toolkit/ippool/request/crawl"
	concPool "github.com/sourcegraph/conc/stream"
	"golang.org/x/exp/slog"
)

func Run(ctx context.Context) error {
	p := concPool.New().WithMaxGoroutines(100)

	crawlers := crawl.Crawlers(ctx)
	for i := range crawlers {
		i := i
		p.Go(func() concPool.Callback {
			ips, err := crawlers[i].Crawl()
			if err != nil {
				global.Logger.ErrorCtx(ctx, "kuaidaili Crawl crawlByPage err", err, slog.Any("page", i))
				return func() {}
			}

			return func() {
				err := pool.Use(ctx).IP.InsertMany(ips)
				if err != nil {
					slog.ErrorCtx(ctx, "Run InsertMany err", err)
				}
			}
		})
	}
	p.Wait()
	return nil
}
