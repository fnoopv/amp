package repository

import (
	"context"

	"github.com/fnoopv/amp/database/model"
	"gorm.io/gorm"
	"goyave.dev/filter"
	"goyave.dev/goyave/v5/database"
	"goyave.dev/goyave/v5/util/errors"
	"goyave.dev/goyave/v5/util/session"
)

// User 用户存储库
type User struct {
	DB *gorm.DB
}

// NewUser 创建新的存储库
func NewUser(db *gorm.DB) *User {
	return &User{
		DB: db,
	}
}

// Paginate 返回分页器
func (us *User) Paginate(ctx context.Context, request *filter.Request) (*database.Paginator[*model.User], error) {
	users := []*model.User{}

	paginator, err := filter.Scope(session.DB(ctx, us.DB), request, &users)

	return paginator, errors.New(err)
}

// FindByID 根据ID获取用户
func (us *User) FindByID(ctx context.Context, id string) (*model.User, error) {
	var user *model.User

	db := us.DB.Where("id", id).First(&user)

	return user, errors.New(db.Error)
}

// FindByUsername 根据用户名获取用户
func (us *User) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	var user *model.User

	db := us.DB.Where("username = ?", username).First(&user)

	return user, errors.New(db.Error)
}

// Delete 根据ID删除用户
func (us *User) Delete(ctx context.Context, id string) error {
	db := us.DB.Delete(&model.User{}, id)

	return errors.New(db.Error)
}

// UpdatePassword 更改密码
func (us *User) UpdatePassword(ctx context.Context, id, password string) error {
	db := us.DB.Model(&model.User{}).Where("id = ?", id).Update("password", password)

	return errors.New(db.Error)
}

// Create 创建用户
func (us *User) Create(ctx context.Context, user *model.User) error {
	db := us.DB.Create(user)
	return errors.New(db.Error)
}

// Update 更新用户信息
func (us *User) Update(ctx context.Context, id string, user *model.User) error {
	db := us.DB.Model(&model.User{}).Where("id = ?", id).Updates(user)

	return errors.New(db.Error)
}
