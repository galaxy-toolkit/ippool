package pool

import (
	"context"

	"github.com/galaxy-toolkit/ippool/domain/model"
	"gorm.io/gorm"
)

// SourceDao 数据访问层
type SourceDao interface {
	GetOneByID(id string) (*model.Source, error)
	GetManyByIDs(ids []string) ([]*model.Source, error)
	InsertOne(i *model.Source) error
	InsertMany(ips []*model.Source) error
	UpdateByID(id string, i *model.Source) (int64, error)

	clone(db *gorm.DB) *sourceDao
}

// NewSourceDao 创建数据访问层
func NewSourceDao(ctx context.Context, db *gorm.DB) SourceDao {
	return &sourceDao{
		ctx: ctx,
		db:  db,
	}
}

type sourceDao struct {
	ctx context.Context
	db  *gorm.DB
}

// GetOneByID 根据 ID 单条查询
func (d *sourceDao) GetOneByID(id string) (*model.Source, error) {
	var data *model.Source

	result := d.db.Where("id = ?", id).First(data)
	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}

// GetManyByIDs 根据 IDs 批量查询
func (d *sourceDao) GetManyByIDs(ids []string) ([]*model.Source, error) {
	return nil, nil
}

// InsertOne 插入一条数据
func (d *sourceDao) InsertOne(i *model.Source) error {
	return nil
}

// InsertMany 插入多条数据
func (d *sourceDao) InsertMany(ips []*model.Source) error {
	return nil
}

// UpdateByID 根据 ID 更新数据
func (d *sourceDao) UpdateByID(id string, i *model.Source) (int64, error) {
	return 0, nil
}

func (d *sourceDao) clone(db *gorm.DB) *sourceDao {
	d.db = d.db.Session(&gorm.Session{Initialized: true}).Session(&gorm.Session{})
	d.db.Statement.ConnPool = db.ConnPool
	return d
}
