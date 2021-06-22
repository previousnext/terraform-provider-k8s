package daemonset

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod"
)

const (
	// FieldName is a field identifier.
	FieldName = "name"
	// FieldNamespace is a field identifier.
	FieldNamespace = "namespace"
	// FieldLabels is a field identifier.
	FieldLabels = "labels"
	// FieldPod is a field identifier.
	FieldPod = "pod"
	// FieldMatchLabels is a field identifier.
	FieldMatchLabels = "match_labels"
)

// Resource returns this packages Resource and Field information.
func Resource() *schema.Resource {
	return &schema.Resource{
		Create: Create,
		Read:   Read,
		Update: Update,
		Delete: Delete,

		Schema: map[string]*schema.Schema{
			FieldName: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			FieldNamespace: &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			FieldLabels: &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			FieldMatchLabels: &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			FieldPod: pod.Fields(),
		},
	}
}
