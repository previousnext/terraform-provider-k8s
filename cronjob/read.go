package cronjob

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

	cronJob, err := conn.BatchV1beta1().CronJobs(namespace).Get(name, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		// This is how we tell Terraform that the resource does not exist.
		d.SetId("")
		return nil
	} else if err != nil {
		return err
	}

	d.Set("name", cronJob.ObjectMeta.Name)
	d.Set("namespace", cronJob.ObjectMeta.Namespace)
	d.Set("labels", cronJob.ObjectMeta.Labels)
	d.Set("schedule", cronJob.Spec.Schedule)
	d.Set("service_account", cronJob.Spec.JobTemplate.Spec.Template.Spec.ServiceAccountName)
	d.Set("container", container.Flatten(cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers))
	d.Set("volume", volume.Flatten(cronJob.Spec.JobTemplate.Spec.Template.Spec.Volumes))

	return nil
}
