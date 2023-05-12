package user

import (
	"context"

	"github.com/galaxy-toolkit/ippool/domain/model"
	"github.com/galaxy-toolkit/ippool/domain/user"
	userModule "github.com/galaxy-toolkit/server/module/user"
	"github.com/galaxy-toolkit/server/server/code"
)

// Service 用户登录 service
type Service struct {
	Ctx context.Context
}

// NewService 创建用户登录 service
func NewService(ctx context.Context) *Service {
	return &Service{
		Ctx: ctx,
	}
}

type LoginParams struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
}

func (s *Service) Login(params *LoginParams) (*model.User, error) {
	u, err := user.Use(s.Ctx).User.FindOne(&user.FindOneOption{Name: params.Username})
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, code.NewError(code.UserNotFound, nil)
	}

	if !userModule.ComparePassword(params.Password, u.Password) {
		return nil, code.NewError(code.PasswordError, nil)
	}

	return u, nil
}
