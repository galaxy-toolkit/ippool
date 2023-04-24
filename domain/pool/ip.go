package pool

import (
	"context"

	"github.com/galaxy-toolkit/ippool/domain/model"
	"gorm.io/gorm"
)

// IPDao 数据访问层
type IPDao interface {
	GetOneByID(id int64) (*model.IP, error)
	GetManyByIDs(ids []int64) ([]*model.IP, error)
	InsertOne(i *model.IP) error
	InsertMany(ips []*model.IP) error
	UpdateByID(id int64, values map[string]any) (int64, error)
	DeleteManyByIDs(ids []int64) (int64, error)

	clone(db *gorm.DB) *ipDao
}

// NewIPDao 创建数据访问层
func NewIPDao(ctx context.Context, db *gorm.DB) IPDao {
	return &ipDao{
		ctx: ctx,
		db:  db,
	}
}

type ipDao struct {
	ctx context.Context
	db  *gorm.DB
}

// GetOneByID 根据 ID 单条查询
func (d *ipDao) GetOneByID(id int64) (*model.IP, error) {
	var data *model.IP

	result := d.db.Where("id = ?", id).First(&data)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}

	return data, nil
}

// GetManyByIDs 根据 IDs 批量查询
func (d *ipDao) GetManyByIDs(ids []int64) ([]*model.IP, error) {
	var data []*model.IP
	result := d.db.Model(&model.IP{}).Where("id IN ?", ids).Find(&data)
	return data, result.Error
}

// InsertOne 插入一条数据
func (d *ipDao) InsertOne(data *model.IP) error {
	result := d.db.Create(data)
	return result.Error
}

// InsertMany 插入多条数据
func (d *ipDao) InsertMany(data []*model.IP) error {
	result := d.db.Create(data)
	return result.Error
}

// UpdateByID 根据 ID 更新数据
func (d *ipDao) UpdateByID(id int64, values map[string]any) (int64, error) {
	result := d.db.Model(&model.IP{}).Where("id = ?", id).UpdateColumns(values)
	return result.RowsAffected, result.Error
}

// DeleteManyByIDs 根据 IDs 批量删除
func (d *ipDao) DeleteManyByIDs(ids []int64) (int64, error) {
	result := d.db.Where("id IN ?", ids).Delete(&model.IP{})
	return result.RowsAffected, result.Error
}

func (d *ipDao) clone(db *gorm.DB) *ipDao {
	d.db = d.db.Session(&gorm.Session{Initialized: true}).Session(&gorm.Session{})
	d.db.Statement.ConnPool = db.ConnPool
	return d
}
