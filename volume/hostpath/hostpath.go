package hostpath

import (
	corev1 "k8s.io/api/core/v1"
)

// Expand will return a structured object.
func Expand(in []interface{}) ([]corev1.Volume, []corev1.VolumeMount, error) {
	if len(in) == 0 {
		return []corev1.Volume{}, []corev1.VolumeMount{}, nil
	}

	var (
		volumes = make([]corev1.Volume, len(in))
		mounts  = make([]corev1.VolumeMount, len(in))
	)

	for i, c := range in {
		p := c.(map[string]interface{})

		if name, ok := p["name"]; ok {
			volumes[i].Name = name.(string)
			mounts[i].Name = name.(string)
		}

		if value, ok := p["source"]; ok {
			volumes[i].HostPath = &corev1.HostPathVolumeSource{
				Path: value.(string),
			}
		}

		if value, ok := p["target"]; ok {
			mounts[i].MountPath = value.(string)
		}
	}

	return volumes, mounts, nil
}

// Flatten structured object into unstructured.
func Flatten(volumes []corev1.Volume, mounts []corev1.VolumeMount) []interface{} {
	att := make([]interface{}, len(volumes))

	for delta, volume := range volumes {
		m := map[string]interface{}{}

		if volume.Name != "" {
			m["name"] = volume.Name

			// Also use this name to lookup the target.
			for _, mount := range mounts {
				if mount.Name == volume.Name {
					m["target"] = mount.MountPath
				}
			}
		}

		if volume.HostPath.Path != "" {
			m["source"] = volume.HostPath.Path
		}

		att[delta] = m
	}

	return att
}
