package serviceaccount

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"

	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Create the ServiceAccount.
func Create(d *schema.ResourceData, m interface{}) error {
	conn := m.(*config.Client)

	serviceaccount, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	out, err := conn.Kubernetes().CoreV1().ServiceAccounts(serviceaccount.ObjectMeta.Namespace).Create(&serviceaccount)
	if err != nil {
		return errors.Wrap(err, "failed to create")
	}

	d.SetId(id.Join(out.ObjectMeta))

	return nil
}
