package hostalias

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/core/v1/pod/container/hostalias/hostname"
)

// Expand will return a structured object.
func Expand(in []interface{}) []corev1.HostAlias {
	if len(in) == 0 {
		return []corev1.HostAlias{}
	}

	aliases := make([]corev1.HostAlias, len(in))

	for key, v := range in {
		value := v.(map[string]interface{})

		if ip, ok := value[FieldIP]; ok && ip != "" {
			aliases[key].IP = ip.(string)
		}

		if hosts, ok := value[hostname.Field]; ok {
			aliases[key].Hostnames = hostname.Expand(hosts.([]interface{}))
		}
	}

	return aliases
}
