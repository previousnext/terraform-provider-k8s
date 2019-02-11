package hostname

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Description: "Hostname(s) to assign to an IP.",
		Optional:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
	}
}
