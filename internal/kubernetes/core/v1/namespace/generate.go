package namespace

import (
	"github.com/hashicorp/terraform/helper/schema"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Generate the Namespace.
func Generate(d *schema.ResourceData) (corev1.Namespace, error) {
	var name = d.Get(FieldName).(string)

	namespace := corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}

	return namespace, nil
}
