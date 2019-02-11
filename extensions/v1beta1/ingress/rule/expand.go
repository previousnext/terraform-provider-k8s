package rule

import (
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"

	"github.com/previousnext/terraform-provider-k8s/extensions/v1beta1/ingress/rule/path"
)

// Expand will return a structured object.
func Expand(in []interface{}) []extensionsv1beta1.IngressRule {
	if len(in) == 0 {
		return []extensionsv1beta1.IngressRule{}
	}

	rules := make([]extensionsv1beta1.IngressRule, len(in))

	for key, v := range in {
		value := v.(map[string]interface{})

		if host, ok := value[FieldHost]; ok && host != "" {
			rules[key].Host = host.(string)
		}

		if paths, ok := value[FieldPath]; ok {
			rules[key].HTTP = &extensionsv1beta1.HTTPIngressRuleValue{
				Paths: path.Expand(paths.([]interface{})),
			}
		}
	}

	return rules
}
