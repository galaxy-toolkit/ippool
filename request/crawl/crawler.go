package crawl

import "github.com/galaxy-toolkit/ippool/domain/model"

// Crawler interface
type Crawler interface {
	Crawl() ([]*model.IP, error)
	CrawlAll() ([]*model.IP, error)
}
