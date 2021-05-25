package storageclass

import (
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	// FieldName is a field identifier.
	FieldName = "name"
	// FieldAnnotations is a field identifier.
	FieldAnnotations = "annotations"
	// FieldLabels is a field identifier.
	FieldLabels = "labels"
	// FieldProvisioner is a field identifier.
	FieldProvisioner = "provisioner"
	// FieldMountOptions is a field identifier.
	FieldMountOptions = "mount_options"
	// FieldParameters is a field identifier.
	FieldParameters = "parameters"
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
			FieldProvisioner: {
				Type:     schema.TypeString,
				Required: true,
			},
			FieldMountOptions: {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			FieldParameters: {
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
	}
}
