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

func (s *sNcuHome) Create(ctx context.Context, ncuHomeInput model.NcuHomeCreateInput) (int64, error) {
	return dao.NcuHome.CreateNcuHome(ctx, &do.NcuHome{
		Sex:   ncuHomeInput.Sex,
		Group: ncuHomeInput.Group,
		Name:  ncuHomeInput.Name,
	})
}

func (s *sNcuHome) Get(ctx context.Context) ([]model.NcuHomeCreateInput, error) {
	var data []model.NcuHomeCreateInput
	err := dao.NcuHome.Ctx(ctx).Scan(&data)
	return data, err
}

func (s *sNcuHome) Update(ctx context.Context, ncuHomeInput model.NcuHomeCreateInput) error {
	_, err := dao.NcuHome.Ctx(ctx).Data(ncuHomeInput).Where("name", ncuHomeInput.Name).Update()
	return err
}
