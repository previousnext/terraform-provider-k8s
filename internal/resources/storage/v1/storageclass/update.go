package storageclass

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
)

// Update the StorageClass.
func Update(d *schema.ResourceData, m interface{}) error {
	conn := m.(*config.Client)

	storageclass := Generate(d)

	_, err := conn.Kubernetes().StorageV1().StorageClasses().Update(&storageclass)
	if err != nil {
		return errors.Wrap(err, "failed to create")
	}

	return nil
}
