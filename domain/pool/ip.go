package pool

import (
	"context"

	"github.com/galaxy-toolkit/ippool/domain/model"
	"github.com/galaxy-toolkit/server/database/dao"
	"gorm.io/gorm"
)

const IPTableSample = " TABLESAMPLE SYSTEM (10)"

// IPDao 数据访问层
type IPDao interface {
	dao.IDao[model.IP]

	Random(size int) ([]*model.IP, error)
}

// NewIPDao 创建数据访问层
func NewIPDao(ctx context.Context, db *gorm.DB) IPDao {
	return &ipDao{
		Dao: dao.Dao[model.IP]{
			Ctx: ctx,
			DB:  db,
		},
	}
}

type ipDao struct {
	dao.Dao[model.IP]
}

// Random 批量随机获取 IP
func (d ipDao) Random(size int) ([]*model.IP, error) {
	data := make([]*model.IP, 0)
	result := d.DB.Table((&model.IP{}).TableName() + IPTableSample).Limit(size).Find(&data)
	return data, result.Error
}
