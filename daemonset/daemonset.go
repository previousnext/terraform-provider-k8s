package daemonset

import (
	"github.com/hashicorp/terraform/helper/schema"

	appsv1beta2 "k8s.io/api/apps/v1beta2"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/container"
	"github.com/previousnext/terraform-provider-k8s/hostaliases"
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
				Description: "Name of the DaemonSet.",
				Required:    true,
			},
			"namespace": {
				Type:        schema.TypeString,
				Description: "Namespace which the DaemonSet will be run in.",
				Required:    true,
			},
			"labels": label.Fields(),
			"host_network": {
				Type:        schema.TypeBool,
				Description: "Run this pod on the hosts network stack.",
				Optional:    true,
				Default:     false,
			},
			"host_pid": {
				Type:        schema.TypeBool,
				Description: "Use the host’s pid namespace.",
				Optional:    true,
			},
			"service_account": {
				Type:        schema.TypeString,
				Description: "ServiceAccount to associate with this DaemonSet.",
				Optional:    true,
			},
			"init_container": container.Fields(),
			"hostaliases":    hostaliases.Fields(),
			"container":      container.Fields(),
			"volume":         volume.Fields(),
		},
	}
}

// Helper function for generating the DaemonSet object.
func generateDaemonSet(d *schema.ResourceData) (appsv1beta2.DaemonSet, error) {
	var (
		name           = d.Get("name").(string)
		namespace      = d.Get("namespace").(string)
		labels         = d.Get("labels").(map[string]interface{})
		hostNetwork    = d.Get("host_network").(bool)
		hostPid        = d.Get("host_pid").(bool)
		serviceAccount = d.Get("service_account").(string)
		initContainer  = d.Get("init_container").([]interface{})
		aliases        = d.Get("hostaliases").([]interface{})
		containers     = d.Get("container").([]interface{})
		volumes        = d.Get("volume").([]interface{})
	)

	initContainerList, err := container.Expand(initContainer)
	if err != nil {
		return appsv1beta2.DaemonSet{}, err
	}

	containerList, err := container.Expand(containers)
	if err != nil {
		return appsv1beta2.DaemonSet{}, err
	}

	volumeList, err := volume.Expand(volumes)
	if err != nil {
		return appsv1beta2.DaemonSet{}, err
	}

	daemonset := appsv1beta2.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    label.Expand(labels),
		},
		Spec: appsv1beta2.DaemonSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: label.Expand(labels),
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: label.Expand(labels),
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: serviceAccount,
					HostNetwork:        hostNetwork,
					InitContainers:     initContainerList,
					Containers:         containerList,
					Volumes:            volumeList,
					HostPID:            hostPid,
					HostAliases:        hostaliases.Expand(aliases),
				},
			},
		},
	}

	return daemonset, nil
}
