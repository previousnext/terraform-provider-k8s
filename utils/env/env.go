package env

import (
	corev1 "k8s.io/api/core/v1"
)

func Expand(in []interface{}) ([]corev1.EnvVar, error) {
	if len(in) == 0 {
		return []corev1.EnvVar{}, nil
	}

	envs := make([]corev1.EnvVar, len(in))

	for i, c := range in {
		p := c.(map[string]interface{})
		if name, ok := p["name"]; ok {
			envs[i].Name = name.(string)
		}
		if value, ok := p["value"]; ok {
			envs[i].Value = value.(string)
		}
	}

	return envs, nil
}

func Flatten(in []corev1.EnvVar) []interface{} {
	att := make([]interface{}, len(in))

	for i, v := range in {
		m := map[string]interface{}{}
		if v.Name != "" {
			m["name"] = v.Name
		}
		if v.Value != "" {
			m["value"] = v.Value
		}
		att[i] = m
	}

	return att
}
