package daemonset

import (
	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/previousnext/terraform-provider-k8s/container"
	"github.com/previousnext/terraform-provider-k8s/utils/id"
	"github.com/previousnext/terraform-provider-k8s/volume"
)

func resourceRead(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	namespace, name, err := id.Split(d.Id())
	if err != nil {
		return err
	}

	daemonset, err := conn.AppsV1beta2().DaemonSets(namespace).Get(name, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		// This is how we tell Terraform that the resource does not exist.
		d.SetId("")
		return nil
	} else if err != nil {
		return err
	}

	d.Set("name", daemonset.ObjectMeta.Name)
	d.Set("namespace", daemonset.ObjectMeta.Namespace)
	d.Set("service_account", daemonset.Spec.Template.Spec.ServiceAccountName)
	d.Set("labels", daemonset.ObjectMeta.Labels)
	d.Set("container", container.Flatten(daemonset.Spec.Template.Spec.Containers))
	d.Set("volume", volume.Flatten(daemonset.Spec.Template.Spec.Volumes))

	return nil
}
