package daemonset

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/apimachinery/objectmeta"
	"github.com/previousnext/terraform-provider-k8s/core/v1/pod"
)

// Generate the DaemonSet.
func Generate(d *schema.ResourceData) (appsv1.DaemonSet, error) {
	var (
		rawMeta = d.Get(objectmeta.FieldObjectMeta).([]interface{})
		rawPod  = d.Get(pod.FieldPod).([]interface{})
	)

	meta := objectmeta.Expand(rawMeta)

	daemonset := appsv1.DaemonSet{
		ObjectMeta: meta,
		Spec: appsv1.DaemonSetSpec{
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
		return daemonset, errors.Wrap(err, "failed to expand Pod")
	}

	daemonset.Spec.Template.Spec = podSpec

	return daemonset, nil
}
