package port

import corev1 "k8s.io/api/core/v1"

// Expand will return a structured object.
func Expand(in []interface{}) ([]corev1.ContainerPort, error) {
	if len(in) == 0 {
		return []corev1.ContainerPort{}, nil
	}

	ports := make([]corev1.ContainerPort, len(in))

	for key, v := range in {
		value := v.(map[string]interface{})

		if name, ok := value[FieldName]; ok {
			ports[key].Name = name.(string)
		}

		if port, ok := value[FieldContainerPort]; ok {
			ports[key].ContainerPort = int32(port.(int))
		}
	}

	return ports, nil
}
