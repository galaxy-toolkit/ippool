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
)

const (
	// CrawlersMaxGoroutines 爬虫最大线程数
	CrawlersMaxGoroutines = 100
)

func CrawlAll(ctx context.Context) {
	resultsChan := make(chan *model.IP, 1000)
	wg := conc.NewWaitGroup()

	wg.Go(func() {
		crawlAll(ctx, resultsChan)
	})
	wg.Go(func() {
		collectResults(ctx, resultsChan)
	})

	wg.Wait()
}

// crawlAll 执行爬虫
func crawlAll(ctx context.Context, resultsChan chan<- *model.IP) {
	p := concPool.New().WithMaxGoroutines(CrawlersMaxGoroutines)

	crawlers := crawl.AllCrawlers(ctx)
	for i := range crawlers {
		i := i
		p.Go(func() {
			err := crawlers[i].CrawlAll(resultsChan)
			if err != nil {
				global.Logger.ErrorCtx(ctx, "kuaidaili CrawlAll crawlByPage err", "err", err)
				return
			}
		})
	}

	p.Wait()
	close(resultsChan)
}

// RunTimedCrawlers 执行定时爬虫
func RunTimedCrawlers(ctx context.Context) {
	resultsChan := make(chan *model.IP, 1000)
	wg := conc.NewWaitGroup()

	wg.Go(func() {
		runTimedCrawlers(ctx, resultsChan)
	})
	wg.Go(func() {
		collectResults(ctx, resultsChan)
	})

	wg.Wait()
}

// runTimedCrawlers 执行定时爬虫
func runTimedCrawlers(ctx context.Context, resultsChan chan<- *model.IP) {
	p := concPool.New().WithMaxGoroutines(CrawlersMaxGoroutines)

	crawlers := crawl.TimedCrawlers(ctx)
	for i := range crawlers {
		i := i
		p.Go(func() {
			err := crawlers[i].TimedCrawl(resultsChan)
			if err != nil {
				global.Logger.ErrorCtx(ctx, "kuaidaili TimedCrawl crawlByPage err", "err", err)
				return
			}
		})
	}

	p.Wait()
	close(resultsChan)
}

// CollectResults 收集爬取结果
func collectResults(ctx context.Context, resultsChan <-chan *model.IP) {
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
