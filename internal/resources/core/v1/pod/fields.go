package pod

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/volume"
)

const (
	// FieldAnnotations is a field identifier.
	FieldAnnotations = "annotations"
	// FieldLabels is a field identifier.
	FieldLabels = "labels"
	// FieldNodeSelector is a field identifier.
	FieldNodeSelector = "node_selector"
	// FieldInitContainer is a field identifier.
	FieldInitContainer = "init_container"
	// FieldContainer is a field identifier.
	FieldContainer = "container"
	// FieldVolume is a field identifier.
	FieldVolume = "volume"
	// FieldServiceAccount is a field identifier.
	FieldServiceAccount = "service_account"
	// FieldPullSecret is a field identifier.
	FieldPullSecret = "pull_secret"
	// FieldHostPID is a field identifier.
	FieldHostPID = "host_pid"
	// FieldPriorityClassName is a field identifier.
	FieldPriorityClassName = "priority_class_name"
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
				FieldAnnotations: {
					Type:     schema.TypeMap,
					Optional: true,
				},
				FieldLabels: {
					Type:     schema.TypeMap,
					Optional: true,
				},
				FieldNodeSelector: {
					Type:     schema.TypeMap,
					Optional: true,
				},
				FieldInitContainer: container.Fields(),
				FieldContainer:     container.Fields(),
				FieldVolume:        volume.Fields(),
				FieldServiceAccount: {
					Type:     schema.TypeString,
					Optional: true,
				},
				FieldPullSecret: {
					Type:     schema.TypeString,
					Optional: true,
				},
				FieldHostPID: {
					Type:        schema.TypeBool,
					Description: "Use the host’s pid namespace.",
					Optional:    true,
				},
				FieldPriorityClassName: {
					Type:        schema.TypeString,
					Description: "Priority indicates the importance of a Pod relative to other Pods.",
					Optional:    true,
				},
			},
		},
	}
}
