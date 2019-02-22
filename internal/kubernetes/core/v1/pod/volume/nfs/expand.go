package nfs

import (
	corev1 "k8s.io/api/core/v1"
)

// Expand will return a structured object.
func Expand(in []interface{}) *corev1.NFSVolumeSource {
	if len(in) == 0 {
		return nil
	}

	source := &corev1.NFSVolumeSource{}

	raw := in[0].(map[string]interface{})

	if val, ok := raw[FieldServer]; ok {
		source.Server = val.(string)
	}

	if val, ok := raw[FieldPath]; ok {
		source.Path = val.(string)
	}

	return source
}
