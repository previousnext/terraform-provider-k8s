package mount

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	// FieldName which is used to identify the name field.
	FieldName = "name"
	// FieldPath which is used to identify the path field.
	FieldPath = "path"
	// FieldReadOnly which is used to identify the readonly field.
	FieldReadOnly = "readonly"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Description: "Mount a volume into a path inside the container.",
		Type:        schema.TypeList,
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				FieldName: {
					Description: "Name of the volume to mount.",
					Type:        schema.TypeString,
					Required:    true,
				},
				FieldPath: {
					Description: "Path to mount the volume into.",
					Type:        schema.TypeString,
					Optional:    true,
				},
				FieldReadOnly: {
					Description: "If this mount is read only.",
					Type:        schema.TypeBool,
					Optional:    true,
				},
			},
		},
	}
}
