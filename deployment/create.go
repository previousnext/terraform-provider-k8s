package deployment

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/client-go/kubernetes"

	"github.com/previousnext/terraform-provider-k8s/utils/id"
)

func resourceCreate(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	deployment, err := generateDeployment(d)
	if err != nil {
		return err
	}

	out, err := conn.AppsV1beta2().Deployments(deployment.ObjectMeta.Namespace).Create(&deployment)
	if err != nil {
		return fmt.Errorf("failed to create apps/v1beta2/deployment: %s", err)
	}

	d.SetId(id.Join(out.ObjectMeta))

	return nil
}
