package envvar

import corev1 "k8s.io/api/core/v1"

// Expand will return a structured object.
func Expand(in []interface{}) []corev1.EnvVar {
	if len(in) == 0 {
		return []corev1.EnvVar{}
	}

	vars := make([]corev1.EnvVar, len(in))

	for key, v := range in {
		value := v.(map[string]interface{})

		if name, ok := value[FieldName]; ok && name != "" {
			vars[key].Name = name.(string)
		}

		if val, ok := value[FieldValue]; ok && val != "" {
			vars[key].Value = val.(string)
		}

		if val, ok := value[FieldValueFieldRef]; ok && val != "" {
			vars[key].ValueFrom = &corev1.EnvVarSource{
				FieldRef: &corev1.ObjectFieldSelector{
					FieldPath: val.(string),
				},
			}
		}
	}

	return vars
}
