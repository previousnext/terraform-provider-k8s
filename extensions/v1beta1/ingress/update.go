package ingress

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
)

// Update the Ingress.
func Update(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	ingress, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	_, err = conn.ExtensionsV1beta1().Ingresses(ingress.ObjectMeta.Namespace).Update(&ingress)
	if err != nil {
		return errors.Wrap(err, "failed to update")
	}

	return nil
}
