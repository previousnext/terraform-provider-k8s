package serviceaccount

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
)

// Update the ServiceAccount.
func Update(d *schema.ResourceData, m interface{}) error {
	conn := m.(*config.Client)

	serviceaccount, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	_, err = conn.Kubernetes().CoreV1().ServiceAccounts(serviceaccount.ObjectMeta.Namespace).Update(&serviceaccount)
	if err != nil {
		return errors.Wrap(err, "failed to update")
	}

	return nil
}
