package nfs

import (
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	// FieldServer is a field identifier.
	FieldServer = "server"
	// FieldPath is a field identifier.
	FieldPath = "path"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				FieldServer: &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				FieldPath: {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	}
}
