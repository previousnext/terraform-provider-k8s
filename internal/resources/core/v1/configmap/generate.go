package configmap

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/interfaceutils"
)

// Generate the Service.
func Generate(d *schema.ResourceData) (corev1.ConfigMap, error) {
	var (
		name      = d.Get(FieldName).(string)
		namespace = d.Get(FieldNamespace).(string)
		rawLabels = d.Get(FieldLabels).(map[string]interface{})
		rawData   = d.Get(FieldData).(map[string]interface{})
	)

	configmap := corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    interfaceutils.ExpandMap(rawLabels),
		},
		Data: interfaceutils.ExpandMap(rawData),
	}

	return configmap, nil
}
