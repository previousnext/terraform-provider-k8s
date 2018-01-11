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
		Elem:        &schema.Schema{Type: schema.TypeString},
	}
}

// Expand will return a structured object.
func Expand(s []interface{}) []string {
	result := make([]string, len(s), len(s))

	for k, v := range s {
		if v == nil {
			// This empty string = "core api" for K8s roles.
			result[k] = ""
		} else {
			result[k] = v.(string)
		}
	}

	return result
}

// Flatten structured object into unstructured.
func Flatten(in []string) []interface{} {
	flattened := make([]interface{}, len(in))

	for _, value := range in {
		flattened = append(flattened, value)
	}

	return flattened
}
