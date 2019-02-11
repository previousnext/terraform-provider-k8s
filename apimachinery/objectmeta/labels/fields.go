package labels

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeMap,
		Description: "Labels to apply to this object so it can be queried",
		Optional:    true,
	}
}
