package user

import (
	"context"
	v1 "github.com/halalala222/gf-learning/api/v1"
	"github.com/halalala222/gf-learning/internal/service"
)

type CUser struct{}

func New() *CUser {
	return &CUser{}
}
func (c *CUser) Login(ctx context.Context, req *v1.LoginReq) (v1.LoginRes, error) {
	res := v1.LoginRes{}
	res.Token, res.Expire = service.Auth().LoginHandler(ctx)
	return res, nil
}
