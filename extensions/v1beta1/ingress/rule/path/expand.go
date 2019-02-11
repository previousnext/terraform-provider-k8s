package path

import (
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// Expand will return a structured object.
func Expand(in []interface{}) []extensionsv1beta1.HTTPIngressPath {
	if len(in) == 0 {
		return []extensionsv1beta1.HTTPIngressPath{}
	}

	paths := make([]extensionsv1beta1.HTTPIngressPath, len(in))

	for key, v := range in {
		value := v.(map[string]interface{})

		if name, ok := value[FieldPath]; ok && name != "" {
			paths[key].Path = name.(string)
		}

		if serviceName, ok := value[FieldService]; ok {
			paths[key].Backend.ServiceName = serviceName.(string)
		}

		if port, ok := value[FieldPort]; ok {
			paths[key].Backend.ServicePort = intstr.FromInt(port.(int))
		}
	}

	return paths
}
