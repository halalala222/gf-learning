package wallet

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/halalala222/gf-learning/internal/consts"
	"github.com/halalala222/gf-learning/internal/dao"
	"github.com/halalala222/gf-learning/internal/model"
	"github.com/halalala222/gf-learning/internal/model/do"
	"github.com/halalala222/gf-learning/internal/service"
	"github.com/halalala222/gf-learning/utility/addressUtils"
)

type sWallet struct{}

func init() {
	service.RegisterWallet(New())
}

func New() service.IWallet {
	return &sWallet{}
}

func (w *sWallet) Login(ctx context.Context, in *model.WalletLoginInput) (*model.WalletLoginOutput, error) {
	in.Address = common.HexToAddress(in.Address).Hex()
	isVerify := addressUtils.Verify(convert(in))
	if !isVerify {
		return nil, gerror.NewCode(consts.ErrInvalidSignature)
	}
	userData, err := dao.User.CreateAndGetOneByAddress(ctx, &do.User{
		Address: in.Address,
	})
	if err != nil {
		log.Error(err.Error())
		return nil, gerror.NewCode(consts.ErrInternal)
	}
	return &model.WalletLoginOutput{
		Id: userData.Id,
	}, nil
}

func convert(in *model.WalletLoginInput) *model.VerifyModel {
	return &model.VerifyModel{
		Address:   in.Address,
		Signature: in.Signature,
		Challenge: in.Challenge,
	}
}
