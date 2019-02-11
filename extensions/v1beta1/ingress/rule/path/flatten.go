package path

import extensionsv1beta1 "k8s.io/api/extensions/v1beta1"

// Flatten structured object into unstructured.
func Flatten(in []extensionsv1beta1.HTTPIngressPath) []interface{} {
	flattened := make([]interface{}, len(in))

	for key, value := range in {
		row := map[string]interface{}{}

		if value.Path != "" {
			row[FieldPath] = value.Path
		}

		if value.Backend.ServiceName != "" {
			row[FieldService] = value.Backend.ServiceName
		}

		if value.Backend.ServicePort.String() != "" {
			row[FieldPort] = value.Backend.ServicePort
		}
		flattened[key] = row
	}

	return flattened
}
