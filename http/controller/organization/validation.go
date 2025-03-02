package organization

import (
	"github.com/fnoopv/amp/http/validation"
	"goyave.dev/goyave/v5"
	v "goyave.dev/goyave/v5/validation"
)

func IndexRequest(request *goyave.Request) v.RuleSet {
	return validation.PaginationRuquest(request)
}

func CreateRequest(_ *goyave.Request) v.RuleSet {
	return v.RuleSet{
		{Path: v.CurrentElement, Rules: v.List{v.Required(), v.Object()}},
	}
}

func UpdateRequest(_ *goyave.Request) v.RuleSet {
	return v.RuleSet{
		{Path: v.CurrentElement, Rules: v.List{v.Required(), v.Object()}},
	}
}
