package storageclass

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	// FieldName is a field identifier.
	FieldName = "name"
	// FieldAnnotations is a field identifier.
	FieldAnnotations = "annotations"
	// FieldLabels is a field identifier.
	FieldLabels = "labels"
	// FieldProvisioner is a field identifier.
	FieldProvisioner = "storage_provisioner"
	// FieldParameters is a field identifier.
	FieldParameters = "parameters"
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
			FieldProvisioner: {
				Type:     schema.TypeString,
				Required: true,
			},
			FieldParameters: &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
	}
}
