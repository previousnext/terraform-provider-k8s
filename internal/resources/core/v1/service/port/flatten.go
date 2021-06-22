package port

import (
	corev1 "k8s.io/api/core/v1"
)

// Flatten structured object into unstructured.
func Flatten(in []corev1.ServicePort) []interface{} {
	flattened := make([]interface{}, len(in))

	for key, value := range in {
		row := map[string]interface{}{}

		if value.Name != "" {
			row[FieldName] = value.Name
		}

		if value.Port > 0 {
			row[FieldPort] = value.Port
		}

		if value.TargetPort.StrVal != "" {
			row[FieldTargetPort] = value.TargetPort.StrVal
		}

		if value.NodePort > 0 {
			row[FieldNodePort] = value.NodePort
		}

		flattened[key] = row
	}

	return flattened
}
