package statefulset

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"

	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Create the Deployment.
func Create(d *schema.ResourceData, m interface{}) error {
	conn := m.(*config.Client)

	statefulset, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	out, err := conn.Kubernetes().AppsV1().StatefulSets(statefulset.ObjectMeta.Namespace).Create(&statefulset)
	if err != nil {
		return errors.Wrap(err, "failed to create")
	}

	d.SetId(id.Join(out.ObjectMeta))

	return nil
}
