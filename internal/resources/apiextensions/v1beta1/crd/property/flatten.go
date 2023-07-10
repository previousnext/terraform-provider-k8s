package property

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// Flatten structured object into unstructured.
func Flatten(in map[string]apiextensionsv1.JSONSchemaProps) []interface{} {
	flattened := make([]interface{}, 0)

	for name, value := range in {
		row := map[string]interface{}{}

		row[FieldName] = name

		if value.Type != "" {
			row[FieldType] = value.Type
		}

		if value.XPreserveUnknownFields != nil {
			row[FieldPreserveUnknownFields] = *value.XPreserveUnknownFields
		}

		flattened = append(flattened, row)
	}

	return flattened
}

