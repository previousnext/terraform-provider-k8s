package subject

import rbacv1 "k8s.io/api/rbac/v1"

// Expand will return a structured object.
func Expand(in []interface{}) []rbacv1.Subject {
	if len(in) == 0 {
		return []rbacv1.Subject{}
	}

	rules := make([]rbacv1.Subject, len(in))

	for key, v := range in {
		value := v.(map[string]interface{})

		if kind, ok := value[FieldKind]; ok {
			rules[key].Kind = kind.(string)
		}

		if name, ok := value[FieldName]; ok {
			rules[key].Name = name.(string)
		}

		if apiGroup, ok := value[FieldAPIGroup]; ok {
			rules[key].APIGroup = apiGroup.(string)
		}

		if namespace, ok := value[FieldNamespace]; ok {
			rules[key].Namespace = namespace.(string)
		}
	}

	return rules
}
