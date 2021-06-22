package clusterrolebinding

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"

	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Create the ClusterRoleBinding.
func Create(d *schema.ResourceData, m interface{}) error {
	conn := m.(*config.Client)

	binding, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	out, err := conn.Kubernetes().RbacV1().ClusterRoleBindings().Create(&binding)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	d.SetId(id.Join(out.ObjectMeta))

	return nil
}
