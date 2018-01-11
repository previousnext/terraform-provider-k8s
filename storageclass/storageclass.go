package storageclass

import (
	"github.com/hashicorp/terraform/helper/schema"
	storagev1beta1 "k8s.io/api/storage/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Resource returns this packages resource.
func Resource() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreate,
		Read:   resourceRead,
		Update: resourceUpdate,
		Delete: resourceDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the StorageClass.",
				Required:    true,
			},
			"storage_provisioner": {
				Type:        schema.TypeString,
				Description: "Provisioner which will be creating PersistentVolumes.",
				Required:    true,
			},
		},
	}
}

// Helper function for generating the StorageClass object.
func generateStorageClass(d *schema.ResourceData) storagev1beta1.StorageClass {
	var (
		name        = d.Get("name").(string)
		provisioner = d.Get("storage_provisioner").(string)
	)

	return storagev1beta1.StorageClass{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Provisioner: provisioner,
	}
}
