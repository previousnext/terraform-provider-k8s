package tcp

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// Expand will return a structured object.
func Expand(in []interface{}) *corev1.TCPSocketAction {
	if len(in) == 0 {
		return nil
	}

	action := &corev1.TCPSocketAction{}

	raw := in[0].(map[string]interface{})

	if port, ok := raw[FieldPort]; ok {
		action.Port = intstr.IntOrString{
			Type:   intstr.String,
			StrVal: port.(string),
		}
	}

	return action
}
