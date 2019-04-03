package hostpath

import (
	corev1 "k8s.io/api/core/v1"
)

// Expand will return a structured object.
func Expand(in []interface{}) *corev1.HostPathVolumeSource {
	if len(in) == 0 {
		return nil
	}

	source := &corev1.HostPathVolumeSource{}

	raw := in[0].(map[string]interface{})

	if val, ok := raw[FieldPath]; ok {
		source.Path = val.(string)
	}

	if val, ok := raw[FieldType]; ok {
		source.Type = val.(*corev1.HostPathType)
	}

	return source
}
