package storageclass

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/previousnext/terraform-provider-k8s/apimachinery/objectmeta"
)

const (
	// FieldProvisioner identifier for the provisioner field.
	FieldProvisioner = "provisioner"
)

// Resource returns this packages Resource and Fields.
func Resource() *schema.Resource {
	return &schema.Resource{
		Create: Create,
		Read:   Read,
		Update: Update,
		Delete: Delete,

		Schema: map[string]*schema.Schema{
			objectmeta.FieldObjectMeta: objectmeta.Fields(),
			FieldProvisioner: {
				Type:        schema.TypeString,
				Description: "Provisioner which will be creating PersistentVolumes.",
				Required:    true,
			},
		},
	}
}
