package hostpath

import (
	corev1 "k8s.io/api/core/v1"
)

// Flatten structured object into unstructured.
func Flatten(in *corev1.HostPathVolumeSource) []interface{} {
	out := make([]interface{}, 1)

	row := map[string]interface{}{}

	if in.Path != "" {
		row[FieldPath] = in.Path
	}
	if *in.Type != corev1.HostPathUnset {
		row[FieldType] = in.Type
	}

	out[0] = row

	return out
}
