package hostpath

import (
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	// FieldPath is a field identifier.
	FieldPath = "path"
	// FieldType is a field identifier.
	FieldType = "type"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				FieldPath: &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				FieldType: &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
			},
		},
	}
}
