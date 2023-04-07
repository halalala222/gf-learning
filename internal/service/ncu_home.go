// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/halalala222/gf-learning/internal/model"
)

type (
	INcuHome interface {
		Create(ctx context.Context, ncuHomeInput model.NcuHomeDto) error
		Get(ctx context.Context) ([]model.NcuHomeDto, error)
	}
)

var (
	localNcuHome INcuHome
)

func NcuHome() INcuHome {
	if localNcuHome == nil {
		panic("implement not found for interface INcuHome, forgot register?")
	}
	return localNcuHome
}

func RegisterNcuHome(i INcuHome) {
	localNcuHome = i
}
