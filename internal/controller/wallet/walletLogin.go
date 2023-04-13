package wallet

import (
	"context"
	v1 "github.com/halalala222/gf-learning/api/v1"
	"github.com/halalala222/gf-learning/internal/service"
)

type CWalletUser struct{}

func New() *CWalletUser {
	return &CWalletUser{}
}
func (c *CWalletUser) Login(ctx context.Context, _ *v1.WalletLoginReq) (res *v1.WalletLoginRes, err error) {
	res = &v1.WalletLoginRes{}
	res.Token, res.Expire = service.Auth().LoginHandler(ctx)
	return
}
