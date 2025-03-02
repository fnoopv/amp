package validation

import (
	"goyave.dev/goyave/v5"
	v "goyave.dev/goyave/v5/validation"
)

// PaginationRuquest 分页验证器
func PaginationRuquest(_ *goyave.Request) v.RuleSet {
	return v.RuleSet{
		{Path: v.CurrentElement, Rules: v.List{v.Required(), v.Object()}},
		{Path: "page", Rules: v.List{v.Required(), v.Int(), v.Min(1)}},
		{Path: "page_size", Rules: v.List{v.Required(), v.Int(), v.In([]int{10, 20, 50, 100})}},
	}
}
