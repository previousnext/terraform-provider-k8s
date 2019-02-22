package daemonset

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
)

// Update the DaemonSet.
func Update(d *schema.ResourceData, m interface{}) error {
	conn := m.(*config.Client)

	daemonset, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	_, err = conn.Kubernetes().AppsV1().DaemonSets(daemonset.ObjectMeta.Namespace).Update(&daemonset)
	if err != nil {
		return errors.Wrap(err, "failed to update")
	}

	return nil
}
