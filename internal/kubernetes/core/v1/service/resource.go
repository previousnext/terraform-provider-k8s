package service

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/core/v1/service/ingress"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/core/v1/service/port"
)

const (
	// FieldName is a field identifier.
	FieldName = "name"
	// FieldNamespace is a field identifier.
	FieldNamespace = "namespace"
	// FieldPort is a field identifier.
	FieldPort = "port"
	// FieldLabels is a field identifier.
	FieldLabels = "labels"
	// FieldSelector is a field identifier.
	FieldSelector = "selector"
	// FieldType is a field identifier.
	FieldType = "type"
	// FieldIngress is a field identifier.
	FieldIngress = "ingress"
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
			FieldNamespace: &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			FieldLabels: &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			FieldType: &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			FieldPort: port.Fields(),
			FieldSelector: &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			FieldIngress: ingress.Fields(),
		},
	}
}
