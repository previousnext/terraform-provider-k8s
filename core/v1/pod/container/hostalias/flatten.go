package hostalias

import corev1 "k8s.io/api/core/v1"

// Flatten structured object into unstructured.
func Flatten(in []corev1.HostAlias) []interface{} {
	flattened := make([]interface{}, len(in))

	for key, value := range in {
		row := map[string]interface{}{}

		if value.IP != "" {
			row["ip"] = value.IP
		}

		if len(value.Hostnames) > 0 {
			row["hostnames"] = value.Hostnames
		}

		flattened[key] = row
	}

	return flattened
}
