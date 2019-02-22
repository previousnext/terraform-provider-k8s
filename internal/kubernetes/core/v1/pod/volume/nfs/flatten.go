package nfs

import (
	corev1 "k8s.io/api/core/v1"
)

// Flatten structured object into unstructured.
func Flatten(in *corev1.NFSVolumeSource) []interface{} {
	out := make([]interface{}, 1)

	row := map[string]interface{}{}

	if in.Server != "" {
		row[FieldServer] = in.Server
	}

	if in.Path != "" {
		row[FieldPath] = in.Path
	}

	out[0] = row

	return out
}
