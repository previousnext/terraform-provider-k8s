package poddisruptionbudget

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	policyv1beta1 "k8s.io/api/policy/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/previousnext/terraform-provider-k8s/internal/interfaceutils"
)

// Generate the Deployment.
func Generate(d *schema.ResourceData) (policyv1beta1.PodDisruptionBudget, error) {
	var (
		name            = d.Get(FieldName).(string)
		namespace       = d.Get(FieldNamespace).(string)
		rawLabels       = d.Get(FieldLabels).(map[string]interface{})
		rawMinAvailable = d.Get(FieldMinAvailable).(string)
		rawMatchLabels  = d.Get(FieldMatchLabels).(map[string]interface{})
	)

	minAvailable := intstr.Parse(rawMinAvailable)

	budget := policyv1beta1.PodDisruptionBudget{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    interfaceutils.ExpandMap(rawLabels),
		},
		Spec: policyv1beta1.PodDisruptionBudgetSpec{
			MinAvailable: &minAvailable,
			Selector: &metav1.LabelSelector{
				MatchLabels: interfaceutils.ExpandMap(rawMatchLabels),
			},
		},
	}

	return budget, nil
}
