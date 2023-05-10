package user

import (
	"context"

	"github.com/galaxy-toolkit/ippool/domain/model"
	"github.com/galaxy-toolkit/server/database/dao"
	"gorm.io/gorm"
)

// Dao 数据访问层
type Dao interface {
	dao.IDao[model.User]
}

// NewUserDao 创建数据访问层
func NewUserDao(ctx context.Context, db *gorm.DB) Dao {
	return &userDao{
		Dao: dao.Dao[model.User]{
			Ctx: ctx,
			DB:  db,
		},
	}
}

type userDao struct {
	dao.Dao[model.User]
}
