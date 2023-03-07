package poddisruptionbudget

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	// FieldName is a field identifier.
	FieldName = "name"
	// FieldNamespace is a field identifier.
	FieldNamespace = "namespace"
	// FieldLabels is a field identifier.
	FieldLabels = "labels"
	// FieldMinAvailable is a field identifier.
	FieldMinAvailable = "min_available"
	// FieldMatchLabels is a field identifier.
	FieldMatchLabels = "match_labels"
)

// Resource returns this packages resource.
func Resource() *schema.Resource {
	return &schema.Resource{
		CreateContext: Create,
		ReadContext:   Read,
		UpdateContext: Update,
		DeleteContext: Delete,

		Schema: map[string]*schema.Schema{
			FieldName: {
				Type:     schema.TypeString,
				Required: true,
			},
			FieldNamespace: {
				Type:     schema.TypeString,
				Optional: true,
			},
			FieldLabels: {
				Type:     schema.TypeMap,
				Optional: true,
			},
			FieldMinAvailable: {
				Type:     schema.TypeString,
				Optional: true,
			},
			FieldMatchLabels: {
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
	}
}
