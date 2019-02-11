package args

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Description: "Args used in conjunction with the command field.",
		Optional:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
	}
}
