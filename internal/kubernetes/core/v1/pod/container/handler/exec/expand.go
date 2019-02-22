package exec

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/core/v1/pod/container/command"
)

// Expand will return a structured object.
func Expand(in []interface{}) *corev1.ExecAction {
	if len(in) == 0 {
		return nil
	}

	action := &corev1.ExecAction{}

	raw := in[0].(map[string]interface{})

	if cmd, ok := raw[FieldCommand]; ok {
		action.Command = command.Expand(cmd.([]interface{}))
	}

	return action
}
