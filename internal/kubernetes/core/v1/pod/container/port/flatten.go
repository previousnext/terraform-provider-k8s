package port

import corev1 "k8s.io/api/core/v1"

// Flatten structured object into unstructured.
func Flatten(in []corev1.ContainerPort) []interface{} {
	flattened := make([]interface{}, len(in))

	for key, value := range in {
		row := map[string]interface{}{}

		if value.Name != "" {
			row[FieldName] = value.Name
		}

		if value.ContainerPort > 0 {
			row[FieldContainerPort] = value.ContainerPort
		}

		flattened[key] = row
	}

	return flattened
}
