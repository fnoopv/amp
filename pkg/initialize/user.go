package initialize

import (
	"reflect"

	"github.com/fnoopv/amp/database/model"
	"github.com/fnoopv/amp/pkg/password"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"goyave.dev/goyave/v5/config"
	"goyave.dev/goyave/v5/util/errors"
)

func init() {
	config.Register("user.email", config.Entry{
		Value:    "",
		Type:     reflect.String,
		IsSlice:  false,
		Required: false,
	},
	)
	config.Register("user.nick_name", config.Entry{
		Value:    "",
		Type:     reflect.String,
		IsSlice:  false,
		Required: true,
	},
	)
	config.Register("user.username", config.Entry{
		Value:    "",
		Type:     reflect.String,
		IsSlice:  false,
		Required: true,
	},
	)
	config.Register("user.password", config.Entry{
		Value:    "",
		Type:     reflect.String,
		IsSlice:  false,
		Required: true,
	},
	)
}

// InitializeUser 系统初始化时添加默认用户
func InitializeUser(db *gorm.DB, email, nick_name, username, pwd string) error {
	var count int64
	if err := db.Model(&model.User{}).Count(&count).Error; err != nil {
		return errors.New(err)
	}
	if count > 0 {
		return nil
	}

	var user model.User

	uid, err := uuid.NewV7()
	if err != nil {
		return errors.New(err)
	}
	user.ID = uid.String()

	hashedPwd, err := password.HashPassword(pwd)
	if err != nil {
		return errors.New(err)
	}
	user.Password = hashedPwd

	user.Email = email
	user.Username = username
	user.NickName = nick_name
	user.Status = "active"
	user.IsMFAActive = false

	err = db.Create(&user).Error

	return errors.New(err)
}
