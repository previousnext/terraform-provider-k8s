package service

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"

	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
)

// Update the ServiceAccount.
func Update(d *schema.ResourceData, m interface{}) error {
	conn := m.(*config.Client)

	service, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	_, err = conn.Kubernetes().CoreV1().Services(service.ObjectMeta.Namespace).Update(&service)
	if err != nil {
		return errors.Wrap(err, "failed to update")
	}

	return nil
}
