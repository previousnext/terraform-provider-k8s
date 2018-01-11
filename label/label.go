package label

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeMap,
		Description: "Labels to apply to his object and its higher level API object",
		Optional:    true,
	}
}

// Expand will return a structured object.
func Expand(in map[string]interface{}) map[string]string {
	labels := make(map[string]string)

	for k, v := range in {
		labels[k] = v.(string)
	}

	return labels
}
