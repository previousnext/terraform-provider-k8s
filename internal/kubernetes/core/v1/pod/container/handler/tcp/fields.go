package tcp

import (
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	// FieldPort which is used to identify the port field.
	FieldPort = "port"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Description: "Perform an action using a TCP action.",
		Type:        schema.TypeList,
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				FieldPort: {
					Description: "Name of the port which will be queried.",
					Type:        schema.TypeString,
					Required:    true,
				},
			},
		},
	}
}
