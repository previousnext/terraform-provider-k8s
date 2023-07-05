package role

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/previousnext/terraform-provider-k8s/internal/resources/rbac/v1/role/rule"
)

const (
	// FieldName is a field identifier.
	FieldName = "name"
	// FieldNamespace is a field identifier.
	FieldNamespace = "namespace"
	// FieldLabels is a field identifier.
	FieldLabels = "labels"
	// FieldRule is a field identifier.
	FieldRule = "rule"
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
			FieldRule: rule.Fields(),
		},
	}
}
