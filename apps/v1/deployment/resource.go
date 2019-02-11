package deployment

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/previousnext/terraform-provider-k8s/apimachinery/objectmeta"
	"github.com/previousnext/terraform-provider-k8s/core/v1/pod"
)

const (
	// FieldReplicas identifies how many replicas a Deployment will rollout.
	FieldReplicas = "replicas"
)

// Resource returns this packages Resource and Fields.
func Resource() *schema.Resource {
	return &schema.Resource{
		Create: Create,
		Read:   Read,
		Update: Update,
		Delete: Delete,

		Schema: map[string]*schema.Schema{
			objectmeta.FieldObjectMeta: objectmeta.Fields(),
			FieldReplicas: {
				Type:        schema.TypeInt,
				Description: "Instances of the application which will be deployed.",
				Required:    true,
				Default:     1,
			},
			pod.FieldPod: pod.Fields(),
		},
	}
}
