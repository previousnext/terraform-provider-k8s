package secret

import (
	"github.com/hashicorp/terraform/helper/schema"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/interfaceutils"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/core/v1/secret/data"
)

// Generate the Service.
func Generate(d *schema.ResourceData) (corev1.Secret, error) {
	var (
		name      = d.Get(FieldName).(string)
		namespace = d.Get(FieldNamespace).(string)
		rawLabels = d.Get(FieldLabels).(map[string]interface{})
		rawType   = d.Get(FieldType).(string)
		rawData   = d.Get(FieldData).(map[string]interface{})
	)

	secret := corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    interfaceutils.ExpandMap(rawLabels),
		},
		Type: corev1.SecretType(rawType),
		Data: data.Expand(rawData),
	}

	return secret, nil
}
