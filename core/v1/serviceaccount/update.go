package serviceaccount

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
)

// Update the ServiceAccount.
func Update(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	serviceaccount, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	_, err = conn.CoreV1().ServiceAccounts(serviceaccount.ObjectMeta.Namespace).Update(&serviceaccount)
	if err != nil {
		return errors.Wrap(err, "failed to update")
	}

	return nil
}
