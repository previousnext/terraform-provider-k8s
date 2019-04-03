package volume

import (
	corev1 "k8s.io/api/core/v1"
)

// Flatten structured object into unstructured.
func Flatten(in []corev1.Volume) []interface{} {
	flattened := make([]interface{}, len(in))

	for key, value := range in {
		row := map[string]interface{}{}

		if value.Name != "" {
			row[FieldName] = value.Name
		}

		if value.PersistentVolumeClaim != nil && value.PersistentVolumeClaim.ClaimName != "" {
			row[FieldPVC] = value.PersistentVolumeClaim.ClaimName
		}

		if value.ConfigMap != nil && value.ConfigMap.LocalObjectReference.Name != "" {
			row[FieldConfigMap] = value.ConfigMap.LocalObjectReference.Name
		}

		if value.HostPath != nil && value.HostPath.Path != "" {
			row[FieldHostPath] = value.HostPath.Path
		}

		if value.EmptyDir != nil {
			row[FieldEmptyDir] = value.EmptyDir.Medium
		}

		flattened[key] = row
	}

	return flattened
}
