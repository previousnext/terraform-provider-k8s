package hostalias

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/core/v1/pod/container/hostalias/hostname"
)

// Flatten structured object into unstructured.
func Flatten(in []corev1.HostAlias) []interface{} {
	flattened := make([]interface{}, len(in))

	for key, value := range in {
		row := map[string]interface{}{}

		if value.IP != "" {
			row[FieldIP] = value.IP
		}

		if len(value.Hostnames) > 0 {
			row[hostname.Field] = value.Hostnames
		}

		flattened[key] = row
	}

	return flattened
}
