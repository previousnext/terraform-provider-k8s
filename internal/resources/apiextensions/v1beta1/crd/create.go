package crd

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"

	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Create the ServiceAccount.
func Create(d *schema.ResourceData, m interface{}) error {
	conn := m.(*config.Client)

	crd, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	out, err := conn.APIExtensions().ApiextensionsV1().CustomResourceDefinitions().Create(&crd)
	if err != nil {
		return errors.Wrap(err, "failed to create")
	}

	d.SetId(id.Join(out.ObjectMeta))

	return nil
}
