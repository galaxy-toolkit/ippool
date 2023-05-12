package user

import (
	"github.com/galaxy-toolkit/ippool/domain/model"
	"github.com/galaxy-toolkit/ippool/domain/user"
	userModule "github.com/galaxy-toolkit/server/module/user"
	"github.com/galaxy-toolkit/server/server/code"
)

type RegisterParams struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (s *Service) Register(params *RegisterParams) error {
	userDao := user.Use(s.Ctx).User

	u, err := userDao.FindOne(&user.FindOneOption{Name: params.Username})
	if err != nil {
		return err
	}
	if u != nil {
		return code.NewError(code.UserExisted, nil)
	}

	err = userDao.InsertOne(&model.User{
		Username: params.Username,
		Password: userModule.MD5(params.Password, userModule.PasswordMD5Salt, userModule.TimesOfPasswordMD5),
	})
	if err != nil {
		return err
	}

	return nil
}
