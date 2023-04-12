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
