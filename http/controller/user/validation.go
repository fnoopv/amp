package user

import (
	"github.com/fnoopv/amp/http/validation"
	"goyave.dev/goyave/v5"
	v "goyave.dev/goyave/v5/validation"
)

func IndexRuquest(request *goyave.Request) v.RuleSet {
	return validation.PaginationRuquest(request)
}

// CreateRequest 创建用户
func CreateRequest(_ *goyave.Request) v.RuleSet {
	return v.RuleSet{
		{Path: v.CurrentElement, Rules: v.List{v.Required(), v.Object()}},
		{Path: "email", Rules: v.List{v.Nullable(), v.Email()}},
		{Path: "nick_name", Rules: v.List{v.Required(), v.String()}},
		{Path: "username", Rules: v.List{v.Required(), v.String()}},
		{Path: "status", Rules: v.List{v.Required(), v.String(), v.In([]string{"active", "inactive"})}},
	}
}

// UpdateRequest 更新用户信息
func UpdateRequest(_ *goyave.Request) v.RuleSet {
	return v.RuleSet{
		{Path: v.CurrentElement, Rules: v.List{v.Required(), v.Object()}},
		{Path: "email", Rules: v.List{v.Nullable(), v.Email()}},
		{Path: "nick_name", Rules: v.List{v.Required(), v.String()}},
		{Path: "username", Rules: v.List{v.Required(), v.String()}},
	}
}

// UpdatePasswordRequest 更改密码
func UpdatePasswordRequest(_ *goyave.Request) v.RuleSet {
	return v.RuleSet{
		{Path: v.CurrentElement, Rules: v.List{v.Required(), v.Object()}},
		{Path: "old_password", Rules: v.List{v.Required(), v.String()}},
		{Path: "new_password", Rules: v.List{v.Required(), v.String()}},
		{Path: "confirm_password", Rules: v.List{v.Required(), v.String(), v.Same("new_password")}},
	}
}
