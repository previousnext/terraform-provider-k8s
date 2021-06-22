package namespace

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
)

// Create the Namespace.
func Create(d *schema.ResourceData, m interface{}) error {
	conn := m.(*config.Client)

	namespace, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	_, err = conn.Kubernetes().CoreV1().Namespaces().Create(&namespace)
	if err != nil {
		return errors.Wrap(err, "failed to create")
	}

	d.SetId(namespace.ObjectMeta.Name)

	return nil
}
