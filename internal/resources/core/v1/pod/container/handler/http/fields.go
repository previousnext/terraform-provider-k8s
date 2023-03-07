package http

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	// FieldPath which is used to identify the path field.
	FieldPath = "path"
	// FieldPort which is used to identify the port field.
	FieldPort = "port"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Description: "Perform an action using a HTTP action.",
		Type:        schema.TypeList,
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				FieldPath: {
					Description: "Path of the http endpoint.",
					Type:        schema.TypeString,
					Required:    true,
				},
				FieldPort: {
					Description: "Name of the port which will be queried.",
					Type:        schema.TypeString,
					Required:    true,
				},
			},
		},
	}
}
