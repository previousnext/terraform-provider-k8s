package ingress

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"

	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Create the Ingress.
func Create(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	ingress, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	out, err := conn.ExtensionsV1beta1().Ingresses(ingress.ObjectMeta.Namespace).Create(&ingress)
	if err != nil {
		return errors.Wrap(err, "failed to create")
	}

	d.SetId(id.Join(out.ObjectMeta))

	return nil
}
