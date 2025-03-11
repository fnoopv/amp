package organization

import (
	"goyave.dev/goyave/v5"
	v "goyave.dev/goyave/v5/validation"
)

func CreateRequest(_ *goyave.Request) v.RuleSet {
	return v.RuleSet{
		{Path: v.CurrentElement, Rules: v.List{v.Required(), v.Object()}},
		{Path: "name", Rules: v.List{v.Required(), v.String()}},
		{Path: "parent_id", Rules: v.List{v.String(), v.Nullable()}},
		{Path: "kind", Rules: v.List{v.Required(), v.String(), v.In([]string{"company", "department"})}},
		{Path: "order", Rules: v.List{v.Int(), v.Nullable()}},
		{Path: "description", Rules: v.List{v.String(), v.Nullable()}},
	}
}

func UpdateRequest(_ *goyave.Request) v.RuleSet {
	return v.RuleSet{
		{Path: v.CurrentElement, Rules: v.List{v.Required(), v.Object()}},
		{Path: "id", Rules: v.List{v.Required(), v.UUID()}},
		{Path: "name", Rules: v.List{v.Required(), v.String()}},
		{Path: "parent_id", Rules: v.List{v.String(), v.Nullable()}},
		{Path: "kind", Rules: v.List{v.Required(), v.String(), v.In([]string{"company", "department"})}},
		{Path: "order", Rules: v.List{v.Int(), v.Nullable()}},
		{Path: "description", Rules: v.List{v.String(), v.Nullable()}},
	}
}

func DeleteRequest(_ *goyave.Request) v.RuleSet {
	return v.RuleSet{
		{Path: v.CurrentElement, Rules: v.List{v.Required(), v.Object()}},
		{Path: "ids", Rules: v.List{v.Required(), v.Array()}},
		{Path: "ids[]", Rules: v.List{v.Required(), v.UUID()}},
	}
}
