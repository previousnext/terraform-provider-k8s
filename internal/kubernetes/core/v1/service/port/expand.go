package port

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// Expand will return a structured object.
func Expand(in []interface{}) []corev1.ServicePort {
	if len(in) == 0 {
		return []corev1.ServicePort{}
	}

	ports := make([]corev1.ServicePort, len(in))

	for key, v := range in {
		value := v.(map[string]interface{})

		if val, ok := value[FieldName]; ok {
			ports[key].Name = val.(string)
		}

		if val, ok := value[FieldPort]; ok {
			ports[key].Port = int32(val.(int))
		}

		if val, ok := value[FieldTargetPort]; ok {
			ports[key].TargetPort = intstr.IntOrString{
				Type:   intstr.String,
				StrVal: val.(string),
			}
		}

		if val, ok := value[FieldNodePort]; ok {
			ports[key].NodePort = int32(val.(int))
		}
	}

	return ports
}
