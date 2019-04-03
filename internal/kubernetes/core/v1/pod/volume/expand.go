package volume

import (
	corev1 "k8s.io/api/core/v1"
)

// Expand will return a structured object.
func Expand(in []interface{}) ([]corev1.Volume, error) {
	if len(in) == 0 {
		return []corev1.Volume{}, nil
	}

	volumes := make([]corev1.Volume, len(in))

	for key, v := range in {
		value := v.(map[string]interface{})

		if name, ok := value[FieldName]; ok {
			volumes[key].Name = name.(string)
		}

		if pvc, ok := value[FieldPVC]; ok && pvc != "" {
			volumes[key].PersistentVolumeClaim = &corev1.PersistentVolumeClaimVolumeSource{
				ClaimName: pvc.(string),
			}
		}

		if cfg, ok := value[FieldConfigMap]; ok && cfg != "" {
			volumes[key].ConfigMap = &corev1.ConfigMapVolumeSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: cfg.(string),
				},
			}
		}

		if hp, ok := value[FieldHostPath]; ok && hp != "" {
			volumes[key].HostPath = &corev1.HostPathVolumeSource{
				Path: hp.(string),
			}
		}

		if dir, ok := value[FieldEmptyDir]; ok && dir != "" {
			volumes[key].EmptyDir = &corev1.EmptyDirVolumeSource{
				Medium: dir.(corev1.StorageMedium),
			}
		}
	}

	return volumes, nil
}
