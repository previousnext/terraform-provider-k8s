package image

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	// FieldName is a field identifier.
	FieldName = "name"
	// FieldNamespace is a field identifier.
	FieldNamespace = "namespace"
	// FieldContainer is a field identifier.
	FieldContainer = "container"
	// FieldFallback is a field identifier.
	FieldFallback = "fallback"
	// FieldResult is a field identifier.
	FieldResult = "result"
)

// Source returns this packages data source.
func Source() *schema.Resource {
	return &schema.Resource{
		ReadContext: Read,

		Schema: map[string]*schema.Schema{
			FieldName: {
				Type:     schema.TypeString,
				Required: true,
			},
			FieldNamespace: {
				Type:     schema.TypeString,
				Required: true,
			},
			FieldContainer: {
				Type:     schema.TypeString,
				Required: true,
			},
			FieldFallback: {
				Type:     schema.TypeString,
				Required: true,
			},
			FieldResult: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
