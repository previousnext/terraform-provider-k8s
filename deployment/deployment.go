package deployment

import (
	"github.com/hashicorp/terraform/helper/schema"

	appsv1beta2 "k8s.io/api/apps/v1beta2"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/container"
	"github.com/previousnext/terraform-provider-k8s/label"
	"github.com/previousnext/terraform-provider-k8s/volume"
)

// Resource returns this packages resource.
func Resource() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreate,
		Read:   resourceRead,
		Update: resourceUpdate,
		Delete: resourceDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the Deployment.",
				Required:    true,
			},
			"namespace": {
				Type:        schema.TypeString,
				Description: "Namespace which the Deployment will be run in.",
				Required:    true,
			},
			"service_account": {
				Type:        schema.TypeString,
				Description: "ServiceAccount to associate with this Deployment.",
				Optional:    true,
			},
			"labels":    label.Fields(),
			"container": container.Fields(),
			"volume":    volume.Fields(),
		},
	}
}

func generateDeployment(d *schema.ResourceData) (appsv1beta2.Deployment, error) {
	var (
		name           = d.Get("name").(string)
		namespace      = d.Get("namespace").(string)
		serviceaccount = d.Get("service_account").(string)
		labels         = d.Get("labels").(map[string]interface{})
		containers     = d.Get("container").([]interface{})
		volumes        = d.Get("volume").([]interface{})
	)

	containerList, err := container.Expand(containers)
	if err != nil {
		return appsv1beta2.Deployment{}, err
	}

	volumeList, err := volume.Expand(volumes)
	if err != nil {
		return appsv1beta2.Deployment{}, err
	}

	deployment := appsv1beta2.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    label.Expand(labels),
		},
		Spec: appsv1beta2.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: label.Expand(labels),
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: label.Expand(labels),
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: serviceaccount,
					Containers:         containerList,
					Volumes:            volumeList,
				},
			},
		},
	}

	return deployment, nil
}
