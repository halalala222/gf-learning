package user

import (
	"context"
	"errors"
	"github.com/halalala222/gf-learning/internal/consts"
	"github.com/halalala222/gf-learning/internal/model"
	"github.com/halalala222/gf-learning/internal/service"
)

type sUser struct {
}

var _ service.IUser = &sUser{}

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

func (s *sUser) Login(ctx context.Context, in model.UserLoginInput) (out *model.UserLoginOutput, err error) {
	if in.Username != consts.FakeUserName || in.Password != consts.FakePassword {
		return &model.UserLoginOutput{}, errors.New("un found user")
	}
	return &model.UserLoginOutput{
		Username: in.Username,
	}, nil
}
