package storageclass

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/client-go/kubernetes"

	"github.com/previousnext/terraform-provider-k8s/utils/id"
)

func resourceCreate(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	storageclass := generateStorageClass(d)

	out, err := conn.StorageV1beta1().StorageClasses().Create(&storageclass)
	if err != nil {
		return fmt.Errorf("failed to create storage/v1beta1/storageclass: %s", err)
	}

	d.SetId(id.Join(out.ObjectMeta))

	return nil
}
