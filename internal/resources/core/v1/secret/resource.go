package secret

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/secret/data"
)

const (
	// FieldName is a field identifier.
	FieldName = "name"
	// FieldNamespace is a field identifier.
	FieldNamespace = "namespace"
	// FieldLabels is a field identifier.
	FieldLabels = "labels"
	// FieldType is a field identifier.
	FieldType = "type"
	// FieldData is a field identifier.
	FieldData = "data"
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
			FieldData: data.Fields(),
		},
	}
}
