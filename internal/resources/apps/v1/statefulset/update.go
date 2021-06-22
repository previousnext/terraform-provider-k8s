package statefulset

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"

	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
)

// Update the Deployment.
func Update(d *schema.ResourceData, m interface{}) error {
	conn := m.(*config.Client)

	statefulset, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	_, err = conn.Kubernetes().AppsV1().StatefulSets(statefulset.ObjectMeta.Namespace).Update(&statefulset)
	if err != nil {
		return errors.Wrap(err, "failed to update")
	}

	return nil
}
