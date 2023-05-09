package pool

import (
	"context"

	"github.com/galaxy-toolkit/ippool/domain/model"
	"github.com/galaxy-toolkit/server/database/dao"
	"gorm.io/gorm"
)

// IPDao 数据访问层
type IPDao interface {
	dao.IDao[model.IP]
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
