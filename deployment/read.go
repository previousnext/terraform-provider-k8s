package deployment

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/previousnext/terraform-provider-k8s/container"
	"github.com/previousnext/terraform-provider-k8s/hostaliases"
	"github.com/previousnext/terraform-provider-k8s/utils/id"
	"github.com/previousnext/terraform-provider-k8s/volume"
)

func resourceRead(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	namespace, name, err := id.Split(d.Id())
	if err != nil {
		return err
	}

	deployment, err := conn.AppsV1beta2().Deployments(namespace).Get(name, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		// This is how we tell Terraform that the resource does not exist.
		d.SetId("")
		return nil
	} else if err != nil {
		return err
	}

	if len(deployment.Spec.Template.Spec.Containers) == 0 {
		return fmt.Errorf("cannot find a container associated with cronjob")
	}

	d.Set("name", deployment.ObjectMeta.Name)
	d.Set("namespace", deployment.ObjectMeta.Namespace)
	d.Set("service_account", deployment.Spec.Template.Spec.ServiceAccountName)
	d.Set("labels", deployment.ObjectMeta.Labels)
	d.Set("image", deployment.Spec.Template.Spec.Containers[0].Image)
	d.Set("hostaliases", hostaliases.Flatten(deployment.Spec.Template.Spec.HostAliases))
	d.Set("container", container.Flatten(deployment.Spec.Template.Spec.Containers))
	d.Set("volume", volume.Flatten(deployment.Spec.Template.Spec.Volumes))

	return nil
}
