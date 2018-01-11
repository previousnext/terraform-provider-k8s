package command

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Description: "Command is the command line to execute inside the container as a part of the job.",
		Optional:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
	}
}

// Expand will return a structured object.
func Expand(s []interface{}) []string {
	result := make([]string, len(s), len(s))

	for k, v := range s {
		result[k] = v.(string)
	}

	return result
}
