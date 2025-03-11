package application

import (
	"goyave.dev/goyave/v5"
	v "goyave.dev/goyave/v5/validation"
)

// CreateRequest 创建应用
func CreateRequest(_ *goyave.Request) v.RuleSet {
	return v.RuleSet{
		{Path: v.CurrentElement, Rules: v.List{v.Required(), v.Object()}},
		{Path: "name", Rules: v.List{v.Required(), v.String()}},
		{Path: "organization_id", Rules: v.List{v.Nullable(), v.String()}},
		{Path: "launched_at", Rules: v.List{v.Nullable(), v.String()}},
		{Path: "description", Rules: v.List{v.String()}},
	}
}

// UpdateRequest 更新应用
func UpdateRequest(_ *goyave.Request) v.RuleSet {
	return v.RuleSet{
		{Path: v.CurrentElement, Rules: v.List{v.Required(), v.Object()}},
		{Path: "id", Rules: v.List{v.Required(), v.UUID()}},
		{Path: "name", Rules: v.List{v.Required(), v.String()}},
		{Path: "organization_id", Rules: v.List{v.Nullable(), v.String()}},
		{Path: "launched_at", Rules: v.List{v.Nullable(), v.String()}},
		{Path: "description", Rules: v.List{v.Nullable(), v.String()}},
	}
}

// DeleteRequest 删除应用
func DeleteRequest(_ *goyave.Request) v.RuleSet {
	return v.RuleSet{
		{Path: v.CurrentElement, Rules: v.List{v.Required(), v.Object()}},
		{Path: "ids", Rules: v.List{v.Required(), v.Array()}},
		{Path: "ids[]", Rules: v.List{v.Required(), v.UUID()}},
	}
}
