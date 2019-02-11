package rule

import (
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"

	"github.com/previousnext/terraform-provider-k8s/extensions/v1beta1/ingress/rule/path"
)

// Flatten structured object into unstructured.
func Flatten(in []extensionsv1beta1.IngressRule) []interface{} {
	flattened := make([]interface{}, len(in))

	for key, value := range in {
		row := map[string]interface{}{}

		if value.Host != "" {
			row[FieldHost] = value.Host
		}

		if value.HTTP != nil {
			row[FieldPath] = path.Flatten(value.HTTP.Paths)
		}

		flattened[key] = row
	}

	return flattened
}
