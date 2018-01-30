package cluster

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/client-go/kubernetes"
)

func resourceUpdate(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	role, err := generateRole(d)
	if err != nil {
		return err
	}

	_, err = conn.RbacV1().ClusterRoles().Update(&role)
	if err != nil {
		return fmt.Errorf("failed to update rbac/v1/role/cluster: %s", err)
	}

	return nil
}
