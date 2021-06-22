package clusterrolebinding

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
)

// Update the ClusterRoleBinding.
func Update(d *schema.ResourceData, m interface{}) error {
	conn := m.(*config.Client)

	binding, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	_, err = conn.Kubernetes().RbacV1().ClusterRoleBindings().Update(&binding)
	if err != nil {
		return errors.Wrap(err, "failed to update")
	}

	return nil
}
