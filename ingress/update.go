package ingress

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/client-go/kubernetes"
)

func resourceUpdate(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	ingress, err := generateIngress(d)
	if err != nil {
		return err
	}

	_, err = conn.ExtensionsV1beta1().Ingresses(ingress.ObjectMeta.Namespace).Update(&ingress)
	if err != nil {
		return fmt.Errorf("failed to update extensions/v1beta1/ingresses: %s", err)
	}

	return nil
}
