package configmap

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"

	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Create the Service.
func Create(d *schema.ResourceData, m interface{}) error {
	conn := m.(*config.Client)

	configmap, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	out, err := conn.Kubernetes().CoreV1().ConfigMaps(configmap.ObjectMeta.Namespace).Create(&configmap)
	if err != nil {
		return errors.Wrap(err, "failed to create")
	}

	d.SetId(id.Join(out.ObjectMeta))

	return nil
}
