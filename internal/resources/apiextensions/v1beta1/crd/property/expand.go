package property

import (
	"k8s.io/utils/pointer"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// Expand will return a structured object.
func Expand(in []interface{}) map[string]apiextensionsv1.JSONSchemaProps {
	if len(in) == 0 {
		return nil
	}

	vars := make(map[string]apiextensionsv1.JSONSchemaProps, len(in))

	for _, v := range in {
		value := v.(map[string]interface{})

		prop := apiextensionsv1.JSONSchemaProps{}

		if val, ok := value[FieldType]; ok && val != "" {
			prop.Type = val.(string)
		}

		if val, ok := value[FieldPreserveUnknownFields]; ok && val == true {
			prop.XPreserveUnknownFields = pointer.Bool(true)
		}

		if val, ok := value[FieldName]; ok && val != "" {
			vars[val.(string)] = prop
		}
	}

	return vars
}

