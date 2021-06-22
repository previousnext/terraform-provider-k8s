package mount

import corev1 "k8s.io/api/core/v1"

// Flatten structured object into unstructured.
func Flatten(in []corev1.VolumeMount) []interface{} {
	flattened := make([]interface{}, len(in))

	for key, value := range in {
		row := map[string]interface{}{}

		if value.Name != "" {
			row[FieldName] = value.Name
		}

		if value.MountPath != "" {
			row[FieldPath] = value.MountPath
		}

		row[FieldReadOnly] = value.ReadOnly

		flattened[key] = row
	}

	return flattened
}
