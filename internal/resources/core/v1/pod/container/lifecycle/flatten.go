package lifecycle

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/handler"
)

// Flatten structured object into unstructured.
func Flatten(in *corev1.Lifecycle) []interface{} {
	out := make([]interface{}, 1)

	row := map[string]interface{}{}

	if in.PostStart != nil {
		row[FieldPostStart] = handler.Flatten(in.PostStart)
	}

	if in.PreStop != nil {
		row[FieldPreStop] = handler.Flatten(in.PreStop)
	}

	out[0] = row

	return out
}
