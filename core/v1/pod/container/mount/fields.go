package mount

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Mount a volume into a path inside the container.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Name of the volume to mount.",
				},
				"path": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Path to mount the volume into.",
				},
				"readonly": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "If this mount is read only.",
				},
			},
		},
	}
}
