package storageclass

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"

	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Create the StorageClass.
func Create(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	storageclass := Generate(d)

	out, err := conn.StorageV1().StorageClasses().Create(&storageclass)
	if err != nil {
		return errors.Wrap(err, "failed to create")
	}

	d.SetId(id.Join(out.ObjectMeta))

	return nil
}
