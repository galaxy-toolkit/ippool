package crawl

import (
	"context"

	"github.com/galaxy-toolkit/ippool/domain/model"
	"github.com/galaxy-toolkit/ippool/request/crawl/kuaidaili"
)

// Crawler interface
type Crawler interface {
	TimedCrawl(chan<- *model.IP) error
	CrawlAll(chan<- *model.IP) error
}

// AllCrawlers 获取所有项目内 IP 网站爬虫
func AllCrawlers(ctx context.Context) []Crawler {
	return []Crawler{
		kuaidaili.New(ctx),
	}
}

// TimedCrawlers 获取项目内需要定时爬取的 IP 网站爬虫
func TimedCrawlers(ctx context.Context) []Crawler {
	return []Crawler{}
}
