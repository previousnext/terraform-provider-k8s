package clusterrolebinding

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
)

// Update the ClusterRoleBinding.
func Update(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	binding, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	_, err = conn.RbacV1().ClusterRoleBindings().Update(&binding)
	if err != nil {
		return errors.Wrap(err, "failed to update")
	}

	return nil
}
