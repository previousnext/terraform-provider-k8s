package serviceaccount

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/previousnext/terraform-provider-k8s/internal/interfaceutils"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Generate the ServiceAccount.
func Generate(d *schema.ResourceData) (corev1.ServiceAccount, error) {
	var (
		name      = d.Get(FieldName).(string)
		namespace = d.Get(FieldNamespace).(string)
		rawLabels = d.Get(FieldLabels).(map[string]interface{})
	)

	serviceaccount := corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    interfaceutils.ExpandMap(rawLabels),
		},
	}

	return serviceaccount, nil
}
