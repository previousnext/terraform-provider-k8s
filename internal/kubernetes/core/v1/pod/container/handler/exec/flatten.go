package exec

import corev1 "k8s.io/api/core/v1"

// Flatten structured object into unstructured.
func Flatten(in *corev1.ExecAction) []interface{} {
	out := make([]interface{}, 1)

	row := map[string]interface{}{}

	if len(in.Command) > 0 {
		row[FieldCommand] = in.Command
	}

	out[0] = row

	return out
}
