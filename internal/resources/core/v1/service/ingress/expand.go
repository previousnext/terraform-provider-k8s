package ingress

import (
	corev1 "k8s.io/api/core/v1"
)

// Expand will return a structured object.
func Expand(in []interface{}) []corev1.LoadBalancerIngress {
	if len(in) == 0 {
		return []corev1.LoadBalancerIngress{}
	}

	ingress := make([]corev1.LoadBalancerIngress, len(in))

	for key, v := range in {
		value := v.(map[string]interface{})

		if val, ok := value[FieldIP]; ok {
			ingress[key].IP = val.(string)
		}

		if val, ok := value[FieldHostname]; ok {
			ingress[key].Hostname = val.(string)
		}
	}

	return ingress
}
