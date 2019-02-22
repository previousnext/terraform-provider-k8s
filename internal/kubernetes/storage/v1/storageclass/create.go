package storageclass

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"

	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Create the StorageClass.
func Create(d *schema.ResourceData, m interface{}) error {
	conn := m.(*config.Client)

	storageclass := Generate(d)

	out, err := conn.Kubernetes().StorageV1().StorageClasses().Create(&storageclass)
	if err != nil {
		return errors.Wrap(err, "failed to create")
	}

	d.SetId(id.Join(out.ObjectMeta))

	return nil
}
