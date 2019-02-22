package namespace

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
)

// Update the Namespace.
func Update(d *schema.ResourceData, m interface{}) error {
	conn := m.(*config.Client)

	namespace, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	_, err = conn.Kubernetes().CoreV1().Namespaces().Update(&namespace)
	if err != nil {
		return errors.Wrap(err, "failed to update")
	}

	return nil
}
