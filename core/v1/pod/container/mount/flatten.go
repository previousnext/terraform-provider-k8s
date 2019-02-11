package mount

import corev1 "k8s.io/api/core/v1"

// Flatten structured object into unstructured.
func Flatten(in []corev1.VolumeMount) []interface{} {
	flattened := make([]interface{}, len(in))

	for key, value := range in {
		row := map[string]interface{}{}

		if value.Name != "" {
			row["name"] = value.Name
		}

		if value.MountPath != "" {
			row["value"] = value.MountPath
		}

		row["readonly"] = value.ReadOnly

		flattened[key] = row
	}

	return flattened
}
