package filling

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
		{Path: "kind_primary", Rules: v.List{v.Required(), v.String(), v.In([]string{"icp", "grade", "public"})}},
		{Path: "kind_secondary", Rules: v.List{v.Nullable(), v.String(), v.In([]string{"", "website", "app", "mini", "quick"})}},
		{Path: "serial_number", Rules: v.List{v.Required(), v.String()}},
		{Path: "completed_at", Rules: v.List{v.Required(), v.String()}},
		{Path: "grade_level", Rules: v.List{v.Nullable(), v.String(), v.In([]string{"", "second", "third", "fourth", "fifth"})}},
		{Path: "proof_attachment_ids", Rules: v.List{v.Nullable(), v.Array()}},
		{Path: "proof_attachment_ids[]", Rules: v.List{v.UUID()}},
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
