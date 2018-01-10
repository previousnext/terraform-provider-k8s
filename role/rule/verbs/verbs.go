package verbs

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Fields() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Description: "Actions which will be accessible eg. get, watch, list etc",
		Optional:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
	}
}

func Expand(s []interface{}) []string {
	result := make([]string, len(s), len(s))

	for k, v := range s {
		result[k] = v.(string)
	}

	return result
}
