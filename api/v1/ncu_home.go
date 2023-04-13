package v1

import "github.com/gogf/gf/v2/frame/g"

type NcuHomeReq struct {
	g.Meta `path:"/" method:"post"`
	Name   string `v:"required"`
	Group  string `v:"required"`
	Sex    string `v:"required"`
}

type NcuHomeRes struct {
	Id int64 `json:"id"`
}
