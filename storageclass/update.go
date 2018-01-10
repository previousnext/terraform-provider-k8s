package storageclass

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/client-go/kubernetes"
)

func resourceUpdate(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	storageclass := generateStorageClass(d)

	_, err := conn.StorageV1beta1().StorageClasses().Update(&storageclass)
	if err != nil {
		return fmt.Errorf("failed to update storage/v1beta2/storageclass: %s", err)
	}

	return nil
}
