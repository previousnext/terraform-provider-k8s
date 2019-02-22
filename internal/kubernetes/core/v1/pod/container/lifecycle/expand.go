package lifecycle

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/core/v1/pod/container/handler"
)

// Expand will return a structured object.
func Expand(in []interface{}) *corev1.Lifecycle {
	if len(in) == 0 {
		return nil
	}

	lifecycle := &corev1.Lifecycle{}

	raw := in[0].(map[string]interface{})

	if post, ok := raw[FieldPostStart]; ok {
		lifecycle.PostStart = handler.Expand(post.([]interface{}))
	}

	if pre, ok := raw[FieldPreStop]; ok {
		lifecycle.PreStop = handler.Expand(pre.([]interface{}))
	}

	return lifecycle
}
