package daemonset

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/client-go/kubernetes"

	"github.com/previousnext/terraform-provider-k8s/utils/id"
)

func resourceCreate(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	daemonset, err := generateDaemonSet(d)
	if err != nil {
		return err
	}

	out, err := conn.AppsV1beta2().DaemonSets(daemonset.ObjectMeta.Namespace).Create(&daemonset)
	if err != nil {
		return fmt.Errorf("failed to create apps/v1beta2/daemonset: %s", err)
	}

	d.SetId(id.Join(out.ObjectMeta))

	return nil
}
