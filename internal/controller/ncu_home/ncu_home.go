package ncu_home

import (
	"context"
	v1 "github.com/halalala222/gf-learning/api/hello/v1"
	"github.com/halalala222/gf-learning/internal/model"
	"github.com/halalala222/gf-learning/internal/service"
)

type cNcuHome struct{}

var CNcuHome = cNcuHome{}

func (c *cNcuHome) CreatePerson(ctx context.Context, req *v1.NcuHomeRep) (v1.NcuHomeRes, error) {
	data := model.NcuHomeDto{
		Sex:   req.Sex,
		Group: req.Group,
		Name:  req.Name,
	}
	lastResultId, err := service.NcuHome().Create(ctx, data)
	return v1.NcuHomeRes{
		Id: lastResultId,
	}, err
}
