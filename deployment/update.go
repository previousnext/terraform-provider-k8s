package deployment

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/client-go/kubernetes"
)

func resourceUpdate(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	deployment, err := generateDeployment(d)
	if err != nil {
		return err
	}

	_, err = conn.AppsV1beta2().Deployments(deployment.ObjectMeta.Namespace).Update(&deployment)
	if err != nil {
		return fmt.Errorf("failed to update apps/v1beta2/deployment: %s", err)
	}

	return nil
}
