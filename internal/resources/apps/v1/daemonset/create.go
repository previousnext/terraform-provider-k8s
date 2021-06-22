package daemonset

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"

	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Create the DaemonSet.
func Create(d *schema.ResourceData, m interface{}) error {
	conn := m.(*config.Client)

	daemonset, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	out, err := conn.Kubernetes().AppsV1().DaemonSets(daemonset.ObjectMeta.Namespace).Create(&daemonset)
	if err != nil {
		return errors.Wrap(err, "failed to create")
	}

	d.SetId(id.Join(out.ObjectMeta))

	return nil
}
