package cronjob

import (
	"github.com/hashicorp/terraform/helper/schema"

	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/container"
	"github.com/previousnext/terraform-provider-k8s/label"
	"github.com/previousnext/terraform-provider-k8s/volume"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreate,
		Read:   resourceRead,
		Update: resourceUpdate,
		Delete: resourceDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the CronJob.",
				Required:    true,
			},
			"namespace": {
				Type:        schema.TypeString,
				Description: "Namespace which the CronJob will be run in.",
				Required:    true,
			},
			"labels": label.Fields(),
			"schedule": {
				Type:        schema.TypeString,
				Description: "How often to run this CronJob.",
				Required:    true,
			},
			"service_account": {
				Type:        schema.TypeString,
				Description: "ServiceAccount to associate with this CronJob.",
				Optional:    true,
			},
			"container": container.Fields(),
			"volume":    volume.Fields(),
		},
	}
}

// Helper function for generating the CronJob object.
func generateCronJob(d *schema.ResourceData) (batchv1beta1.CronJob, error) {
	var (
		name           = d.Get("name").(string)
		namespace      = d.Get("namespace").(string)
		labels         = d.Get("labels").(map[string]interface{})
		schedule       = d.Get("schedule").(string)
		serviceaccount = d.Get("service_account").(string)
		containers     = d.Get("container").([]interface{})
		volumes        = d.Get("volume").([]interface{})
	)

	containerList, err := container.Expand(containers)
	if err != nil {
		return batchv1beta1.CronJob{}, err
	}

	volumeList, err := volume.Expand(volumes)
	if err != nil {
		return batchv1beta1.CronJob{}, err
	}

	cronJob := batchv1beta1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    label.Expand(labels),
		},
		Spec: batchv1beta1.CronJobSpec{
			Schedule: schedule,
			JobTemplate: batchv1beta1.JobTemplateSpec{
				Spec: batchv1.JobSpec{
					Template: corev1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Labels: label.Expand(labels),
						},
						Spec: corev1.PodSpec{
							RestartPolicy:      corev1.RestartPolicyNever,
							Containers:         containerList,
							Volumes:            volumeList,
							ServiceAccountName: serviceaccount,
						},
					},
				},
			},
		},
	}

	return cronJob, nil
}
