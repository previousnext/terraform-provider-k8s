package clusterrole

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
)

// Update the ClusterRole.
func Update(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	role, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	_, err = conn.RbacV1().ClusterRoles().Update(&role)
	if err != nil {
		return errors.Wrap(err, "failed to update")
	}

	return nil
}
