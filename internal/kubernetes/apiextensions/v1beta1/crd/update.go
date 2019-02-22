package crd

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
)

// Update the StorageClass.
func Update(d *schema.ResourceData, m interface{}) error {
	conn := m.(*config.Client)

	crd, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	_, err = conn.APIExtensions().ApiextensionsV1beta1().CustomResourceDefinitions().Update(&crd)
	if err != nil {
		return errors.Wrap(err, "failed to update")
	}

	return nil
}
