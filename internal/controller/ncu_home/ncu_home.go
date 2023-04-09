package ncu_home

import (
	"context"
	"github.com/halalala222/gf-learning/api/v1"
	"github.com/halalala222/gf-learning/internal/model"
	"github.com/halalala222/gf-learning/internal/service"
)

type CNcuHome struct{}

func New() *CNcuHome {
	return &CNcuHome{}
}

func (c *CNcuHome) CreatePerson(ctx context.Context, req *v1.NcuHomeReq) (v1.NcuHomeRes, error) {
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
