package lifecycle

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/core/v1/pod/container/handler"
)

const (
	// FieldPostStart is a field identifier.
	FieldPostStart = "post_start"
	// FieldPreStop is a field identifier.
	FieldPreStop = "pre_stop"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Description: "Lifecycle events for a container.",
		Type:        schema.TypeList,
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				FieldPostStart: handler.Fields(),
				FieldPreStop:   handler.Fields(),
			},
		},
	}
}
