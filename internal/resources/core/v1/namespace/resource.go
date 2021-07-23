package namespace

import (
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	// FieldName is a field identifier.
	FieldName = "name"
	// FieldLabels is a field identifier.
	FieldLabels = "labels"
)

// Resource returns this packages Resource and Fields.
func Resource() *schema.Resource {
	return &schema.Resource{
		Create: Create,
		Read:   Read,
		Update: Update,
		Delete: Delete,

		Schema: map[string]*schema.Schema{
			FieldName: {
				Type:     schema.TypeString,
				Required: true,
			},
			FieldLabels: {
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
	}
}
