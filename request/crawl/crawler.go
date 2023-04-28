package crawl

import (
	"context"

	"github.com/galaxy-toolkit/ippool/domain/model"
	"github.com/galaxy-toolkit/ippool/request/crawl/kuaidaili"
)

// Crawler interface
type Crawler interface {
	Crawl(chan<- *model.IP) error
	CrawlAll(chan<- *model.IP) error
}

// Crawlers 获取项目内 IP 网站爬虫
func Crawlers(ctx context.Context) []Crawler {
	return []Crawler{
		kuaidaili.New(ctx),
	}
}
