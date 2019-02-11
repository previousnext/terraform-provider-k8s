package deployment

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"

	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Create the Deployment.
func Create(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	deployment, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	out, err := conn.AppsV1().Deployments(deployment.ObjectMeta.Namespace).Create(&deployment)
	if err != nil {
		return errors.Wrap(err, "failed to create")
	}

	d.SetId(id.Join(out.ObjectMeta))

	return nil
}
