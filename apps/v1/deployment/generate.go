package deployment

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/apimachinery/objectmeta"
	"github.com/previousnext/terraform-provider-k8s/core/v1/pod"
)

// Generate the Deployment.
func Generate(d *schema.ResourceData) (appsv1.Deployment, error) {
	var (
		rawMeta  = d.Get(objectmeta.FieldObjectMeta).([]interface{})
		replicas = d.Get(FieldReplicas).(int32)
		rawPod   = d.Get(pod.FieldPod).([]interface{})
	)

	meta := objectmeta.Expand(rawMeta)

	deployment := appsv1.Deployment{
		ObjectMeta: meta,
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: meta.Labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: meta.Labels,
				},
			},
		},
	}

	podSpec, err := pod.Expand(rawPod)
	if err != nil {
		return deployment, errors.Wrap(err, "failed to expand Pod")
	}

	deployment.Spec.Template.Spec = podSpec

	return deployment, nil
}
