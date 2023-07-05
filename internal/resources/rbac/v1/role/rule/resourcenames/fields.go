package resourcenames

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Description: "Name of resources which will be accessible eg. pods",
		Optional:    true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
}
