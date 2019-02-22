package http

import corev1 "k8s.io/api/core/v1"

// Flatten structured object into unstructured.
func Flatten(in *corev1.HTTPGetAction) []interface{} {
	out := make([]interface{}, 1)

	row := map[string]interface{}{}

	if in.Port.StrVal != "" {
		row[FieldPort] = in.Port.StrVal
	}

	if in.Path != "" {
		row[FieldPath] = in.Path
	}

	out[0] = row

	return out
}
