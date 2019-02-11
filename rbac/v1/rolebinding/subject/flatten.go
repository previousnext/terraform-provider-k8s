package subject

import rbacv1 "k8s.io/api/rbac/v1"

// Flatten structured object into unstructured.
func Flatten(in []rbacv1.Subject) []interface{} {
	flattened := make([]interface{}, len(in))

	for key, value := range in {
		row := map[string]interface{}{}

		if value.Kind != "" {
			row[FieldKind] = value.Kind
		}

		if value.Name != "" {
			row[FieldName] = value.Name
		}

		if value.APIGroup != "" {
			row[FieldAPIGroup] = value.APIGroup
		}

		if value.Namespace != "" {
			row[FieldNamespace] = value.Namespace
		}

		flattened[key] = row
	}

	return flattened
}
