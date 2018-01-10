package serviceaccount

import (
	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/previousnext/terraform-provider-k8s/utils/id"
)

func resourceRead(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	namespace, name, err := id.Split(d.Id())
	if err != nil {
		return err
	}

	serviceaccount, err := conn.CoreV1().ServiceAccounts(namespace).Get(name, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		// This is how we tell Terraform that the resource does not exist.
		d.SetId("")
		return nil
	} else if err != nil {
		return err
	}

	d.Set("name", serviceaccount.ObjectMeta.Name)
	d.Set("namespace", serviceaccount.ObjectMeta.Namespace)

	return nil
}
