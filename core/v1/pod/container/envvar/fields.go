package envvar

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Description: "Environment variables which can be set for a container",
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Name of environment variable.",
				},
				"value": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Environment variable value.",
				},
				"field_path": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Field path to use for environment variable.",
				},
			},
		},
	}
}
