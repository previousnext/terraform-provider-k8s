package daemonset

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/client-go/kubernetes"
)

func resourceUpdate(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	daemonset, err := generateDaemonSet(d)
	if err != nil {
		return err
	}

	_, err = conn.AppsV1beta2().DaemonSets(daemonset.ObjectMeta.Namespace).Update(&daemonset)
	if err != nil {
		return fmt.Errorf("failed to update apps/v1beta2/daemonset: %s", err)
	}

	return nil
}
