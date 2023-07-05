package configmap

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
	// FieldData is a field identifier.
	FieldData = "data"
)

// Resource returns this packages resource.
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
			FieldData: &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
	}
}
