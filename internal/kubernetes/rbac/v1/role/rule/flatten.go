package rule

import rbacv1 "k8s.io/api/rbac/v1"

// Flatten structured object into unstructured.
func Flatten(in []rbacv1.PolicyRule) []interface{} {
	flattened := make([]interface{}, len(in))

	for key, value := range in {
		row := map[string]interface{}{}

		if len(value.APIGroups) > 0 {
			row[FieldAPIGroups] = value.APIGroups
		}

		if len(value.Resources) > 0 {
			row[FieldResources] = value.Resources
		}

		if len(value.ResourceNames) > 0 {
			row[FieldResourceNames] = value.ResourceNames
		}

		if len(value.Verbs) > 0 {
			row[FieldVerbs] = value.Verbs
		}

		flattened[key] = row
	}

	return flattened
}
