package deployment

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/interfaceutils"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/core/v1/pod"
)

// Generate the Deployment.
func Generate(d *schema.ResourceData) (appsv1.Deployment, error) {
	var (
		name           = d.Get(FieldName).(string)
		namespace      = d.Get(FieldNamespace).(string)
		rawLabels      = d.Get(FieldLabels).(map[string]interface{})
		replicas       = int32(d.Get(FieldReplicas).(int))
		rawMatchLabels = d.Get(FieldMatchLabels).(map[string]interface{})
		rawPod         = d.Get(FieldPod).([]interface{})
	)

	deployment := appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    interfaceutils.ExpandMap(rawLabels),
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: interfaceutils.ExpandMap(rawMatchLabels),
			},
		},
	}

	template, err := pod.Expand(rawPod)
	if err != nil {
		return deployment, errors.Wrap(err, "failed to expand Pod")
	}

	deployment.Spec.Template = template

	return deployment, nil
}
