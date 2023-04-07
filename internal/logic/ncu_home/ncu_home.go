package ncu_home

import (
	"context"
	"github.com/halalala222/gf-learning/internal/dao"
	"github.com/halalala222/gf-learning/internal/model"
	"github.com/halalala222/gf-learning/internal/model/do"
	"github.com/halalala222/gf-learning/internal/service"
)

type (
	sNcuHome struct{}
)

func init() {
	service.RegisterNcuHome(New())
}

func New() service.INcuHome {
	return &sNcuHome{}
}

func (s *sNcuHome) Create(ctx context.Context, ncuHomeInput model.NcuHomeDto) error {
	_, err := dao.NcuHome.Ctx(ctx).Data(do.NcuHome{
		Sex:   ncuHomeInput.Sex,
		Name:  ncuHomeInput.Name,
		Group: ncuHomeInput.Group,
	}).Insert()
	return err
}

func (s *sNcuHome) Get(ctx context.Context) ([]model.NcuHomeDto, error) {
	var data []model.NcuHomeDto
	err := dao.NcuHome.Ctx(ctx).Scan(&data)
	return data, err
}

func (s *sNcuHome) Update(ctx context.Context, ncuHomeInput model.NcuHomeDto) error {
	_, err := dao.NcuHome.Ctx(ctx).Data(ncuHomeInput).Where("name", ncuHomeInput.Name).Update()
	return err
}
