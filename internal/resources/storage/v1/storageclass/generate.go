package storageclass

import (
	"github.com/hashicorp/terraform/helper/schema"
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/interfaceutils"
)

// Generate the StorageClass.
func Generate(d *schema.ResourceData) storagev1.StorageClass {
	var (
		name        = d.Get(FieldName).(string)
		rawLabels   = d.Get(FieldLabels).(map[string]interface{})
		provisioner = d.Get(FieldProvisioner).(string)
		parameters  = d.Get(FieldParameters).(map[string]interface{})
	)

	return storagev1.StorageClass{
		ObjectMeta: metav1.ObjectMeta{
			Name:   name,
			Labels: interfaceutils.ExpandMap(rawLabels),
		},
		Provisioner: provisioner,
		Parameters:  interfaceutils.ExpandMap(parameters),
	}
}
