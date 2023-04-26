package kuaidaili

import (
	"context"
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/galaxy-toolkit/ippool/domain/model"
	"github.com/wnanbei/fastreq"
)

const (
	pageURL = "https://www.kuaidaili.com/free/inha/%s"
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
	uri := fmt.Sprintf(pageURL, "1")
	resp, err := fastreq.Get(uri)
	if err != nil {
		return nil, err
	}

	if err := resp.BuildDom(); err != nil {
		return nil, err
	}

	ips := make([]*model.IP, 0)
	resp.Dom.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
		ip := &model.IP{
			Address:   s.Find(`td[data-title="IP"]`).Text(),
			Port:      s.Find(`td[data-title="PORT"]`).Text(),
			Protocol:  s.Find(`td[data-title="类型"]`).Text(),
			Location:  s.Find(`td[data-title="位置"]`).Text(),
			Source:    source,
			CrawlTime: time.Now().UnixMilli(),
		}
		ips = append(ips, ip)
	})

	for i := range ips {
		fmt.Println(*ips[i])
	}
	return ips, nil
}
