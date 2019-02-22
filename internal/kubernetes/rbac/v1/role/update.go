package role

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
)

// Update the Role.
func Update(d *schema.ResourceData, m interface{}) error {
	conn := m.(*config.Client)

	role, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	_, err = conn.Kubernetes().RbacV1().Roles(role.ObjectMeta.Namespace).Update(&role)
	if err != nil {
		return errors.Wrap(err, "failed to update")
	}

	return nil
}
