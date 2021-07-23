package namespace

import (
	"github.com/hashicorp/terraform/helper/schema"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/interfaceutils"
)

// Generate the Namespace.
func Generate(d *schema.ResourceData) (corev1.Namespace, error) {
	var (
		name   = d.Get(FieldName).(string)
		labels = d.Get(FieldLabels).(map[string]interface{})
	)

	namespace := corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:   name,
			Labels: interfaceutils.ExpandMap(labels),
		},
	}

	return namespace, nil
}
