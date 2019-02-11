package rolebinding

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
)

// Update the RoleBinding.
func Update(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	binding, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	_, err = conn.RbacV1().RoleBindings(binding.ObjectMeta.Namespace).Update(&binding)
	if err != nil {
		return errors.Wrap(err, "failed to update")
	}

	return nil
}
