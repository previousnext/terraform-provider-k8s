package exec

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/core/v1/pod/container/command"
)

const (
	// FieldCommand which is used to identify the command field.
	FieldCommand = "command"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Description: "Perform an action using a exec action.",
		Type:        schema.TypeList,
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				FieldCommand: command.Fields(),
			},
		},
	}
}
