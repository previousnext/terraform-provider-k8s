package ingress

import (
	corev1 "k8s.io/api/core/v1"
)

// Flatten structured object into unstructured.
func Flatten(in []corev1.LoadBalancerIngress) []interface{} {
	flattened := make([]interface{}, len(in))

	for key, value := range in {
		row := map[string]interface{}{}

		if value.IP != "" {
			row[FieldIP] = value.IP
		}

		if value.Hostname != "" {
			row[FieldHostname] = value.Hostname
		}

		flattened[key] = row
	}

	return flattened
}
