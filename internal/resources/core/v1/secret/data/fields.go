package data

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeMap,
		Optional: true,
	}
}
