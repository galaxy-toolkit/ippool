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

	FindOne(opt *FindOneOption) (*model.User, error)
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

// FindOneOption 查询用户选项
type FindOneOption struct {
	Name  string
	Phone string
	Email string
}

// FindOne 根据条件查询单个用户
func (d userDao) FindOne(opt *FindOneOption) (*model.User, error) {
	if opt == nil {
		return nil, nil
	}

	query := d.DB
	if opt.Name != "" {
		query = query.Where("username = ?", opt.Name)
	}
	if opt.Phone != "" {
		query = query.Where("phone = ?", opt.Phone)
	}
	if opt.Email != "" {
		query = query.Where("email = ?", opt.Email)
	}

	var user model.User
	result := query.First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}
