package ingress

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/client-go/kubernetes"

	"github.com/previousnext/terraform-provider-k8s/utils/id"
)

func resourceCreate(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	ingress, err := generateIngress(d)
	if err != nil {
		return err
	}

	out, err := conn.ExtensionsV1beta1().Ingresses(ingress.ObjectMeta.Namespace).Create(&ingress)
	if err != nil {
		return fmt.Errorf("failed to create extensions/v1beta1/ingresses: %s", err)
	}

	d.SetId(id.Join(out.ObjectMeta))

	return nil
}
