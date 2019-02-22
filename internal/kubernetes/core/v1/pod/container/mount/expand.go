package mount

import corev1 "k8s.io/api/core/v1"

// Expand will return a structured object.
func Expand(in []interface{}) ([]corev1.VolumeMount, error) {
	if len(in) == 0 {
		return []corev1.VolumeMount{}, nil
	}

	mounts := make([]corev1.VolumeMount, len(in))

	for key, v := range in {
		value := v.(map[string]interface{})

		if name, ok := value[FieldName]; ok {
			mounts[key].Name = name.(string)
		}

		if path, ok := value[FieldPath]; ok {
			mounts[key].MountPath = path.(string)
		}

		if readonly, ok := value[FieldReadOnly]; ok {
			mounts[key].ReadOnly = readonly.(bool)
		}
	}

	return mounts, nil
}
