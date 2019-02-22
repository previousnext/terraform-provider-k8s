package names

import (
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

// Flatten structured object into unstructured.
func Flatten(in apiextensionsv1beta1.CustomResourceDefinitionNames) []interface{} {
	out := make([]interface{}, 1)

	row := map[string]interface{}{}

	if in.Plural != "" {
		row[FieldPlural] = in.Plural
	}

	if in.Singular != "" {
		row[FieldSingular] = in.Singular
	}

	if len(in.ShortNames) > 0 {
		row[FieldShortNames] = in.ShortNames
	}

	if in.Kind != "" {
		row[FieldKind] = in.Kind
	}

	out[0] = row

	return out
}
