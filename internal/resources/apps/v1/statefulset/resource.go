package statefulset

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

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
	// FieldReplicas is a field identifier.
	FieldReplicas = "replicas"
	// FieldMatchLabels is a field identifier.
	FieldMatchLabels = "match_labels"
)

// Resource returns this packages Resource and Fields.
func Resource() *schema.Resource {
	return &schema.Resource{
		CreateContext: Create,
		ReadContext:   Read,
		UpdateContext: Update,
		DeleteContext: Delete,

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
			FieldReplicas: {
				Type:     schema.TypeInt,
				Required: true,
			},
			FieldMatchLabels: &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			FieldPod: pod.Fields(),
		},
	}
}
