package storageclass

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
)

// Update the StorageClass.
func Update(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	storageclass := Generate(d)

	_, err := conn.StorageV1().StorageClasses().Update(&storageclass)
	if err != nil {
		return errors.Wrap(err, "failed to create")
	}

	return nil
}
