package apigroups

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Description: "APIs which will be accessible. Empty means that it will be the core API.",
		Optional:    true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
}
