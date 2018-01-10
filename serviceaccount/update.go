package serviceaccount

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/client-go/kubernetes"
)

func resourceUpdate(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	serviceaccount, err := generateServiceAccount(d)
	if err != nil {
		return err
	}

	_, err = conn.CoreV1().ServiceAccounts(serviceaccount.ObjectMeta.Namespace).Update(&serviceaccount)
	if err != nil {
		return fmt.Errorf("failed to update core/v1/serviceaccount: %s", err)
	}

	return nil
}
