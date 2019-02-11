package clusterrolebinding

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"

	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Create the ClusterRoleBinding.
func Create(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	binding, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	out, err := conn.RbacV1().ClusterRoleBindings().Create(&binding)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	d.SetId(id.Join(out.ObjectMeta))

	return nil
}
