package envvar

import (
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	// FieldName is a field identifier.
	FieldName = "name"
	// FieldValue is a field identifier.
	FieldValue = "value"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Description: "Environment variables which can be set for a container",
		Type:        schema.TypeList,
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				FieldName: {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Name of environment variable.",
				},
				FieldValue: {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Environment variable value.",
				},
			},
		},
	}
}
