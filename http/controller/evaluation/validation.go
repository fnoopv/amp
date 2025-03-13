package evaluation

import (
	"goyave.dev/goyave/v5"
	v "goyave.dev/goyave/v5/validation"
)

// CreateRequest 创建
func CreateRequest(_ *goyave.Request) v.RuleSet {
	return v.RuleSet{
		{Path: v.CurrentElement, Rules: v.List{v.Required(), v.Object()}},
		{Path: "filling_id", Rules: v.List{v.Nullable(), v.String()}},
		{Path: "completed_at", Rules: v.List{v.Required(), v.String()}},
		{Path: "serial_number", Rules: v.List{v.Required(), v.String()}},
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
