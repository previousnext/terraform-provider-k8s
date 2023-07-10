package property

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	// FieldName is a field identifier.
	FieldName = "name"
	// FieldType is a field identifier.
	FieldType = "type"
	// FieldPreserveUnknownFields is a field identifier.
	FieldPreserveUnknownFields = "preserve_unknown_fields"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				FieldName: {
					Type:     schema.TypeString,
					Required: true,
				},
				FieldType: {
					Type:     schema.TypeString,
					Required: true,
				},
				FieldPreserveUnknownFields: {
					Type:     schema.TypeBool,
					Optional: true,
				},
			},
		},
	}
}
