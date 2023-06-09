// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/halalala222/gf-learning/internal/dao/internal"
	"github.com/halalala222/gf-learning/internal/model/do"
	"github.com/halalala222/gf-learning/internal/model/entity"
)

// internalUserDao is internal type for wrapping internal DAO implements.
type internalUserDao = *internal.UserDao

// userDao is the data access object for table user.
// You can define custom methods on it to extend its functionality as you wish.
type userDao struct {
	internalUserDao
}

var (
	// User is globally public accessible object for table user operations.
	User = userDao{
		internal.NewUserDao(),
	}
)

// Fill with you ideas below.

func (u *userDao) GetOneByAddress(ctx context.Context, user *do.User) (*entity.User, bool, error) {
	record, err := User.Ctx(ctx).Where("address = ?", user.Address).One()
	if err != nil {
		return nil, false, err
	}
	if record.IsEmpty() {
		return nil, false, nil
	}
	userData := &entity.User{}
	err = record.Struct(userData)
	return userData, true, nil
}

func (u *userDao) CreateAndGetId(ctx context.Context, user *do.User) (*entity.User, error) {
	id, err := User.Ctx(ctx).Data(&do.User{
		Address:   user.Address,
		CreatedAt: gtime.Now(),
		DeletedAt: nil,
		UpdatedAt: nil,
	}).InsertAndGetId()
	return &entity.User{
		Id: int(id),
	}, err
}

func (u *userDao) CreateAndGetOneByAddress(ctx context.Context, user *do.User) (*entity.User, error) {
	userData, isExited, err := u.GetOneByAddress(ctx, user)
	if err != nil {
		return nil, err
	}
	if !isExited {
		userData, err = u.CreateAndGetId(ctx, user)
	}
	return userData, err
}
