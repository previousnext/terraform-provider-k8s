package resource

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "CPU and Memory constraints to apply.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"cpu": {
					Type:        schema.TypeString,
					Description: "CPU resource constraint.",
					Optional:    true,
				},
				"memory": {
					Type:        schema.TypeString,
					Description: "Memory resource constraint.",
					Optional:    true,
				},
			},
		},
	}
}
