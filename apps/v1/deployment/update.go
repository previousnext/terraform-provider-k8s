package deployment

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
)

// Update the Deployment.
func Update(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	deployment, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	_, err = conn.AppsV1().Deployments(deployment.ObjectMeta.Namespace).Update(&deployment)
	if err != nil {
		return errors.Wrap(err, "failed to update")
	}

	return nil
}
