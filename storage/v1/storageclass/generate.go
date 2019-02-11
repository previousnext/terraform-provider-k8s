package storageclass

import (
	"github.com/hashicorp/terraform/helper/schema"
	storagev1 "k8s.io/api/storage/v1"

	"github.com/previousnext/terraform-provider-k8s/apimachinery/objectmeta"
)

// Generate the StorageClass.
func Generate(d *schema.ResourceData) storagev1.StorageClass {
	var (
		rawMeta     = d.Get(objectmeta.FieldObjectMeta).([]interface{})
		provisioner = d.Get(FieldProvisioner).(string)
	)

	return storagev1.StorageClass{
		ObjectMeta:  objectmeta.Expand(rawMeta),
		Provisioner: provisioner,
	}
}
