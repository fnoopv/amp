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

// Organization 组织存储库
type Organization struct {
	DB *gorm.DB
}

// NewOrganization 创建新的存储库
func NewOrganization(db *gorm.DB) *Organization {
	return &Organization{
		DB: db,
	}
}

// Paginate 返回分页器
func (us *Organization) Paginate(ctx context.Context, request *filter.Request) (*database.Paginator[*model.Organization], error) {
	users := []*model.Organization{}

	paginator, err := filter.Scope(session.DB(ctx, us.DB), request, &users)

	return paginator, errors.New(err)
}

// FindByID 根据ID获取组织信息
func (us *Organization) FindByID(ctx context.Context, id string) (*model.Organization, error) {
	var org *model.Organization

	db := us.DB.Where("id", id).First(&org)

	return org, errors.New(db.Error)
}

// Delete 根据ID列表删除组织
func (us *Organization) Delete(ctx context.Context, id string) error {
	db := us.DB.Delete(&model.Organization{}, id)

	return errors.New(db.Error)
}

// Create 创建组织
func (us *Organization) Create(ctx context.Context, user *model.Organization) error {
	return us.DB.Create(user).Error
}

// Update 更新组织信息
func (us *Organization) Update(ctx context.Context, id string, organization *model.Organization) error {
	return us.DB.Model(&model.Organization{}).Where("id = ?", id).Updates(organization).Error
}
