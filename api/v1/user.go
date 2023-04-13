package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type LoginReq struct {
	g.Meta `path:"/login" method:"POST"`
}

type LoginRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type WalletLoginReq struct {
	g.Meta    `path:"/wallet/login" method:"POST"`
	Signature string `v:"required" json:"signature"`
	Address   string `v:"required" json:"address"`
	Challenge string `v:"required" json:"challenge"`
}

type WalletLoginRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}
