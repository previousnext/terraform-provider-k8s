package service

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/core/v1/service/ingress"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/core/v1/service/port"
)

// Source returns this packages data sources.
func Source() *schema.Resource {
	return &schema.Resource{
		Read: Read,

		Schema: map[string]*schema.Schema{
			FieldName: &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the object.",
				Required:    true,
			},
			FieldNamespace: &schema.Schema{
				Type:        schema.TypeString,
				Description: "Namespace which the object will reside.",
				Optional:    true,
			},
			FieldLabels: &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			FieldType: &schema.Schema{
				Type:        schema.TypeString,
				Description: "Type determines how the Service is exposed.",
				Optional:    true,
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
