package pool

import (
	"context"
	"github.com/galaxy-toolkit/ippool/domain/model"
	"github.com/galaxy-toolkit/ippool/domain/pool"
	"github.com/galaxy-toolkit/ippool/internal/global"
	"github.com/galaxy-toolkit/ippool/request/crawl"
	"github.com/jackc/pgconn"
	"github.com/sourcegraph/conc"
	concPool "github.com/sourcegraph/conc/pool"
	"golang.org/x/exp/slog"
)

const (
	// CrawlersMaxGoroutines 爬虫最大线程数
	CrawlersMaxGoroutines = 100
)

func Run(ctx context.Context) {
	resultsChan := make(chan *model.IP, 1000)
	wg := conc.NewWaitGroup()

	wg.Go(func() {
		RunCrawlers(ctx, resultsChan)
	})
	wg.Go(func() {
		CollectResults(ctx, resultsChan)
	})

	wg.Wait()
}

// RunCrawlers 执行爬虫
func RunCrawlers(ctx context.Context, resultsChan chan<- *model.IP) {
	p := concPool.New().WithMaxGoroutines(CrawlersMaxGoroutines)

	crawlers := crawl.Crawlers(ctx)
	for i := range crawlers {
		i := i
		p.Go(func() {
			err := crawlers[i].Crawl(resultsChan)
			if err != nil {
				global.Logger.ErrorCtx(ctx, "kuaidaili CrawlAll crawlByPage err", "err", err, slog.Any("page", i))
				return
			}
		})
	}

	p.Wait()
	close(resultsChan)
}

// CollectResults 收集爬取结果
func CollectResults(ctx context.Context, resultsChan <-chan *model.IP) {
	for i := range resultsChan {
		err := pool.Use(ctx).IP.InsertOne(i)
		if err != nil {
			if e, ok := err.(*pgconn.PgError); ok { // 当错误类型为唯一键冲突时，仅记录 warning 日志
				if e.Code == "23505" {
					global.Logger.WarnCtx(ctx, "CollectResults InsertOne duplicated", "err", err)
					continue
				}
			}

			global.Logger.ErrorCtx(ctx, "CollectResults InsertOne err", "err", err)
		}
	}
}
