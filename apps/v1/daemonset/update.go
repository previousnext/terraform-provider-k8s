package daemonset

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
)

// Update the DaemonSet.
func Update(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	daemonset, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	_, err = conn.AppsV1().DaemonSets(daemonset.ObjectMeta.Namespace).Update(&daemonset)
	if err != nil {
		return errors.Wrap(err, "failed to update")
	}

	return nil
}
