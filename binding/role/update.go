package role

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/client-go/kubernetes"
)

func resourceUpdate(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	binding, err := generateRoleBinding(d)
	if err != nil {
		return err
	}

	_, err = conn.RbacV1().RoleBindings(binding.ObjectMeta.Namespace).Update(&binding)
	if err != nil {
		return fmt.Errorf("failed to update rbac/v1/rolebinding: %s", err)
	}

	return nil
}
