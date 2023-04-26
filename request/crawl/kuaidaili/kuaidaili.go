package kuaidaili

import (
	"context"
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/galaxy-toolkit/ippool/domain/model"
	"github.com/galaxy-toolkit/ippool/internal/global"
	"github.com/wnanbei/fastreq"
	"golang.org/x/exp/slog"
)

const (
	DailyCrawlPage = 3    // 每日定时爬取页数
	MaxCrawlPage   = 1000 // 最大爬取数量
)

const (
	pageURL = "https://www.kuaidaili.com/free/inha/%d"
	source  = "www.kuaidaili.com"
)

type Crawler struct {
	ctx context.Context
}

func New(ctx context.Context) *Crawler {
	return &Crawler{
		ctx: ctx,
	}
}

func (c *Crawler) Crawl() ([]*model.IP, error) {
	ips := make([]*model.IP, 0, DailyCrawlPage*10)

	for i := 1; i <= DailyCrawlPage; i++ {
		ip, err := c.crawlByPage(i)
		if err != nil {
			global.Logger.ErrorCtx(c.ctx, "kuaidaili Crawl crawlByPage err", err, slog.Any("page", i))
			return nil, err
		}
		ips = append(ips, ip...)
	}

	return ips, nil
}

func (c *Crawler) CrawlAll() ([]*model.IP, error) {
	ips := make([]*model.IP, 0, DailyCrawlPage*10)

	for i := 1; i <= MaxCrawlPage; i++ {
		ip, err := c.crawlByPage(i)
		if err != nil {
			global.Logger.ErrorCtx(c.ctx, "kuaidaili CrawlAll crawlByPage err", err, slog.Any("page", i))
			return nil, err
		}
		ips = append(ips, ip...)
	}

	return ips, nil
}

// crawlByPage 分页爬取
func (c *Crawler) crawlByPage(page int) ([]*model.IP, error) {
	uri := fmt.Sprintf(pageURL, page)
	resp, err := fastreq.Get(uri)
	if err != nil {
		return nil, err
	}

	dom, err := resp.Dom()
	if err != nil {
		return nil, err
	}

	ips := make([]*model.IP, 0)
	dom.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
		ip := &model.IP{
			Address:   s.Find(`td[data-title="IP"]`).Text(),
			Port:      s.Find(`td[data-title="PORT"]`).Text(),
			Protocol:  model.IPProtocol(s.Find(`td[data-title="类型"]`).Text()),
			Location:  s.Find(`td[data-title="位置"]`).Text(),
			Source:    source,
			CrawlTime: time.Now().UnixMilli(),
		}
		ips = append(ips, ip)
	})

	return ips, nil
}
