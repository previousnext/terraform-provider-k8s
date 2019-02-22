package names

import (
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	// FieldPlural is a field identifier.
	FieldPlural = "plural"
	// FieldSingular is a field identifier.
	FieldSingular = "singular"
	// FieldShortNames is a field identifier.
	FieldShortNames = "short_names"
	// FieldKind is a field identifier.
	FieldKind = "kind"
)

// Fields which define a Pod.
func Fields() *schema.Schema {
	return &schema.Schema{
		Description: ".",
		Type:        schema.TypeList,
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				FieldPlural: {
					Description: ".",
					Type:        schema.TypeString,
					Required:    true,
				},
				FieldSingular: {
					Description: ".",
					Type:        schema.TypeString,
					Optional:    true,
				},
				FieldShortNames: {
					Description: ".",
					Type:        schema.TypeList,
					Optional:    true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
				FieldKind: {
					Description: ".",
					Type:        schema.TypeString,
					Required:    true,
				},
			},
		},
	}
}
