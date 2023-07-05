package names

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/interfaceutils"
)

// Expand will return a structured object.
func Expand(in []interface{}) apiextensionsv1.CustomResourceDefinitionNames {
	var names apiextensionsv1.CustomResourceDefinitionNames

	if len(in) == 0 {
		return names
	}

	raw := in[0].(map[string]interface{})

	if val, ok := raw[FieldPlural]; ok {
		names.Plural = val.(string)
	}

	if val, ok := raw[FieldSingular]; ok {
		names.Singular = val.(string)
	}

	if val, ok := raw[FieldShortNames]; ok {
		names.ShortNames = interfaceutils.ExpandSlice(val.([]interface{}))
	}

	if val, ok := raw[FieldKind]; ok {
		names.Kind = val.(string)
	}

	return names
}
