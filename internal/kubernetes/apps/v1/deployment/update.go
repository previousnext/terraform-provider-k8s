package deployment

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"

	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
)

// Update the Deployment.
func Update(d *schema.ResourceData, m interface{}) error {
	conn := m.(*config.Client)

	deployment, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	_, err = conn.Kubernetes().AppsV1().Deployments(deployment.ObjectMeta.Namespace).Update(&deployment)
	if err != nil {
		return errors.Wrap(err, "failed to update")
	}

	return nil
}
