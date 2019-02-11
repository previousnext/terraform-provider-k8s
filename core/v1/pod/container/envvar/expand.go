package envvar

import corev1 "k8s.io/api/core/v1"

// Expand will return a structured object.
func Expand(in []interface{}) []corev1.EnvVar {
	if len(in) == 0 {
		return []corev1.EnvVar{}
	}

	mounts := make([]corev1.EnvVar, len(in))

	for key, v := range in {
		value := v.(map[string]interface{})

		if name, ok := value["name"]; ok && name != "" {
			mounts[key].Name = name.(string)
		}

		if val, ok := value["value"]; ok && val != "" {
			mounts[key].Value = val.(string)
		}

		if path, ok := value["field_path"]; ok && path != "" {
			mounts[key].ValueFrom = &corev1.EnvVarSource{
				FieldRef: &corev1.ObjectFieldSelector{
					FieldPath: path.(string),
				},
			}
		}
	}

	return mounts
}
