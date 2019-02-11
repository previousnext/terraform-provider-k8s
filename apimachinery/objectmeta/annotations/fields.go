package annotations

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeMap,
		Description: "Annotations to apply to this object for additional context",
		Optional:    true,
	}
}
