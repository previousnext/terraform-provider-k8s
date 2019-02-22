package clusterrole

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/rbac/v1/role/rule"
)

const (
	// FieldName is a field identifier.
	FieldName = "name"
	// FieldLabels is a field identifier.
	FieldLabels = "labels"
	// FieldRule is a field identifier.
	FieldRule = "rule"
)

// Resource returns this packages resource.
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
			FieldLabels: &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			FieldRule: rule.Fields(),
		},
	}
}
