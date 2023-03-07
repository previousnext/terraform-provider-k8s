package crd

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/previousnext/terraform-provider-k8s/internal/resources/apiextensions/v1beta1/crd/names"
)

const (
	// FieldName is a field identifier.
	FieldName = "name"
	// FieldAnnotations is a field identifier.
	FieldAnnotations = "annotations"
	// FieldLabels is a field identifier.
	FieldLabels = "labels"
	// FieldGroup is a field identifier.
	FieldGroup = "group"
	// FieldVersion is a field identifier.
	FieldVersion = "version"
	// FieldScope is a field identifier.
	FieldScope = "scope"
	// FieldNames is a field identifier.
	FieldNames = "names"
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
				Type:     schema.TypeString,
				Required: true,
			},
			FieldLabels: &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			FieldGroup: {
				Type:     schema.TypeString,
				Required: true,
			},
			FieldVersion: {
				Type:     schema.TypeString,
				Required: true,
			},
			FieldScope: {
				Type:     schema.TypeString,
				Required: true,
			},
			FieldNames: names.Fields(),
		},
	}
}
