package ip

import (
	"context"

	"github.com/galaxy-toolkit/ippool/domain/model/ip"
)

// Dao 数据访问层
type Dao interface {
	GetOneByID(id string) (*ip.IP, error)
	GetManyByIDs(ids []string) ([]*ip.IP, error)
	InsertOne(i *ip.IP) error
	InsertMany(ips []*ip.IP) error
	UpdateByID(id string, i *ip.IP) (int64, error)
}

type dao struct {
	ctx context.Context
}

// GetOneByID 根据 ID 单条查询
func (d dao) GetOneByID(id string) (*ip.IP, error) {
	return nil, nil
}

// GetManyByIDs 根据 IDs 批量查询
func (d dao) GetManyByIDs(ids []string) ([]*ip.IP, error) {
	return nil, nil
}

// InsertOne 插入一条数据
func (d dao) InsertOne(i *ip.IP) error {
	return nil
}

// InsertMany 插入多条数据
func (d dao) InsertMany(ips []*ip.IP) error {
	return nil
}

// UpdateByID 根据 ID 更新数据
func (d dao) UpdateByID(id string, i *ip.IP) (int64, error) {
	return 0, nil
}
