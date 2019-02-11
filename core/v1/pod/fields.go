package pod

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/previousnext/terraform-provider-k8s/core/v1/pod/container"
	"github.com/previousnext/terraform-provider-k8s/core/v1/pod/volume"
)

const (
	// FieldPod identifier for Pod objects.
	FieldPod            = "pod"
	FieldInitContainer  = "init_container"
	FieldContainer      = "container"
	FieldVolume         = "volume"
	FieldHostAlias      = "host_alias"
	FieldServiceAccount = "service_account"
)

// Fields which define a Pod.
func Fields() *schema.Schema {
	return &schema.Schema{
		Description: "List of Volumes used to mount into containers.",
		Type:        schema.TypeList,
		Required:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				FieldInitContainer: container.Fields(),
				FieldContainer:     container.Fields(),
				FieldVolume:        volume.Fields(),
				FieldServiceAccount: {
					Type:        schema.TypeString,
					Description: "ServiceAccount to associate with this Pod.",
					Optional:    true,
				},
			},
		},
	}
}
