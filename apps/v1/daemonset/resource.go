package daemonset

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/previousnext/terraform-provider-k8s/apimachinery/objectmeta"
	"github.com/previousnext/terraform-provider-k8s/core/v1/pod"
)

// Resource returns this packages Resource and Field information.
func Resource() *schema.Resource {
	return &schema.Resource{
		Create: Create,
		Read:   Read,
		Update: Update,
		Delete: Delete,

		Schema: map[string]*schema.Schema{
			objectmeta.FieldObjectMeta: objectmeta.Fields(),
			pod.FieldPod:               pod.Fields(),
		},
	}
}
