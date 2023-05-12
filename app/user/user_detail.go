package user

import (
	"github.com/galaxy-toolkit/ippool/domain/model"
	"github.com/galaxy-toolkit/ippool/domain/user"
	"github.com/galaxy-toolkit/server/server/code"
)

func (s *Service) Detail(uid int64) (*model.User, error) {
	u, err := user.Use(s.Ctx).User.GetOneByID(uid)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, code.NewError(code.UserNotFound, nil)
	}

	return u, nil
}
