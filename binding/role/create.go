package role

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/client-go/kubernetes"

	"github.com/previousnext/terraform-provider-k8s/utils/id"
)

func resourceCreate(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	binding, err := generateRoleBinding(d)
	if err != nil {
		return err
	}

	out, err := conn.RbacV1().RoleBindings(binding.ObjectMeta.Namespace).Create(&binding)
	if err != nil {
		return fmt.Errorf("failed to create rbac/v1/rolebinding: %s", err)
	}

	d.SetId(id.Join(out.ObjectMeta))

	return nil
}
