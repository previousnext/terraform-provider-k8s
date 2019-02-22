package resource

import (
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	// FieldCPU is used to identify the cpu field.
	FieldCPU = "cpu"
	// FieldMemory is used to identify the memory field.
	FieldMemory = "memory"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Description: "CPU and Memory constraints to apply.",
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				FieldCPU: {
					Description: "CPU resource constraint.",
					Type:        schema.TypeString,
					Optional:    true,
				},
				FieldMemory: {
					Description: "Memory resource constraint.",
					Type:        schema.TypeString,
					Optional:    true,
				},
			},
		},
	}
}
