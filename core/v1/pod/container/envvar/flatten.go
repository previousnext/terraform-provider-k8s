package envvar

import corev1 "k8s.io/api/core/v1"

// Flatten structured object into unstructured.
func Flatten(in []corev1.EnvVar) []interface{} {
	flattened := make([]interface{}, len(in))

	for key, value := range in {
		row := map[string]interface{}{}

		if value.Name != "" {
			row["name"] = value.Name
		}

		if value.Value != "" {
			row["value"] = value.Value
		}

		if value.ValueFrom != nil && value.ValueFrom.FieldRef != nil && value.ValueFrom.FieldRef.FieldPath != "" {
			row["field_path"] = value.ValueFrom.FieldRef.FieldPath
		}

		flattened[key] = row
	}

	return flattened
}
