package daemonset

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/interfaceutils"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/core/v1/pod"
)

// Generate the DaemonSet.
func Generate(d *schema.ResourceData) (appsv1.DaemonSet, error) {
	var (
		name           = d.Get(FieldName).(string)
		namespace      = d.Get(FieldNamespace).(string)
		rawLabels      = d.Get(FieldLabels).(map[string]interface{})
		rawMatchLabels = d.Get(FieldMatchLabels).(map[string]interface{})
		rawPod         = d.Get(FieldPod).([]interface{})
	)

	daemonset := appsv1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    interfaceutils.ExpandMap(rawLabels),
		},
		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: interfaceutils.ExpandMap(rawMatchLabels),
			},
		},
	}

	template, err := pod.Expand(rawPod)
	if err != nil {
		return daemonset, errors.Wrap(err, "failed to expand Pod")
	}

	daemonset.Spec.Template = template

	return daemonset, nil
}
