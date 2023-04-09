package v1

import "github.com/gogf/gf/v2/frame/g"

type NcuHomeRep struct {
	g.Meta `path:"/ncuhome" method:"post"`
	Name   string `v:"required"`
	Group  string `v:"required"`
	Sex    string `v:"required"`
}

type NcuHomeRes struct {
	Id int64 `json:"id"`
}
