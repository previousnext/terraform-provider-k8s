package command

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Description: "Command is the command line to execute inside the container as a part of the job.",
		Type:        schema.TypeList,
		Optional:    true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
}
