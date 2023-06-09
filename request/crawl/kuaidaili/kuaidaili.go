package kuaidaili

import (
	"context"
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/galaxy-toolkit/ippool/domain/model"
	"github.com/galaxy-toolkit/server/log"
	"github.com/wnanbei/fastreq"
	"golang.org/x/exp/slog"
)

const (
	TimedCrawlPage    = 3                // 定时爬取页数
	MaxCrawlPage      = 500              // 最大爬取数量
	CrawlIntervalTime = time.Second * 10 // 每次爬取间隔时间
)

const (
	pageURL = "https://www.kuaidaili.com/free/inha/%d"
	source  = "www.kuaidaili.com"
)

type Crawler struct {
	ctx                    context.Context
	timedCrawlPage         int
	timedCrawlIntervalTime time.Duration
	maxCrawlPage           int
	crawlIntervalTime      time.Duration
}

func New(ctx context.Context) *Crawler {
	return &Crawler{
		ctx:               ctx,
		timedCrawlPage:    TimedCrawlPage,
		maxCrawlPage:      MaxCrawlPage,
		crawlIntervalTime: CrawlIntervalTime,
	}
}

func (c *Crawler) TimedCrawl(resultChan chan<- *model.IP) error {
	for i := 1; i <= c.timedCrawlPage; i++ {
		ips, err := c.crawlByPage(i)
		if err != nil {
			log.Basic.Error(c.ctx, "kuaidaili Crawl crawlByPage err", slog.Any("err", err), slog.Any("page", i))
			return err
		}

		for j := range ips {
			resultChan <- ips[j]
			log.Basic.Info(c.ctx, "kuaidaili Crawl successes", slog.Any("ip", ips[j]))
		}

		time.Sleep(c.crawlIntervalTime)
	}

	return nil
}

func (c *Crawler) CrawlAll(resultChan chan<- *model.IP) error {
	for i := 1; i <= c.maxCrawlPage; i++ {
		ips, err := c.crawlByPage(i * 10)
		if err != nil {
			log.Basic.Error(c.ctx, "kuaidaili CrawlAll crawlByPage err", slog.Any("err", err), slog.Any("page", i))
			return err
		}

		for j := range ips {
			resultChan <- ips[j]
			log.Basic.Info(c.ctx, "kuaidaili CrawlAll successes", slog.Any("ip", ips[j]))
		}

		time.Sleep(c.crawlIntervalTime)
	}

	return nil
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
			Address:  s.Find(`td[data-title="IP"]`).Text(),
			Port:     s.Find(`td[data-title="PORT"]`).Text(),
			Protocol: model.IPProtocol(s.Find(`td[data-title="类型"]`).Text()),
			Location: s.Find(`td[data-title="位置"]`).Text(),
			Source:   source,
			Status:   model.NotVerify,
			CreateAt: time.Now(),
		}
		ips = append(ips, ip)
	})

	log.Basic.Info(c.ctx, "kuaidaili crawlByPage successes", slog.Any("url", uri))
	return ips, nil
}
