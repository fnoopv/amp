package network

import (
	"goyave.dev/goyave/v5"
	v "goyave.dev/goyave/v5/validation"
)

// CreateRequest 创建
func CreateRequest(_ *goyave.Request) v.RuleSet {
	return v.RuleSet{
		{Path: v.CurrentElement, Rules: v.List{v.Required(), v.Object()}},
		{Path: "name", Rules: v.List{v.Required(), v.String()}},
		{Path: "organization_id", Rules: v.List{v.Nullable(), v.String()}},
		{Path: "filling_id", Rules: v.List{v.Nullable(), v.UUID()}},
		{Path: "description", Rules: v.List{v.Nullable(), v.String()}},
	}
}

// UpdateRequest 更新
func UpdateRequest(request *goyave.Request) v.RuleSet {
	return v.RuleSet{
		{Path: v.CurrentElement, Rules: v.List{v.Required(), v.Object()}},
		{Path: "id", Rules: v.List{v.Required(), v.UUID()}},
		{Path: v.CurrentElement, Rules: CreateRequest(request)},
	}
}

// DeleteRequest 删除
func DeleteRequest(_ *goyave.Request) v.RuleSet {
	return v.RuleSet{
		{Path: v.CurrentElement, Rules: v.List{v.Required(), v.Object()}},
		{Path: "ids", Rules: v.List{v.Required(), v.Array()}},
		{Path: "ids[]", Rules: v.List{v.Required(), v.UUID()}},
	}
}
