package rule

import (
	"github.com/hashicorp/terraform/helper/schema"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"

	"github.com/previousnext/terraform-provider-k8s/ingress/rule/path"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Description: "Rules to apply to an Ingress",
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"host": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Hostname to respond to.",
				},
				"path": path.Fields(),
			},
		},
	}
}

// Expand will return a structured object.
func Expand(in []interface{}) []extensionsv1beta1.IngressRule {
	if len(in) == 0 {
		return []extensionsv1beta1.IngressRule{}
	}

	rules := make([]extensionsv1beta1.IngressRule, len(in))

	for key, v := range in {
		value := v.(map[string]interface{})

		if host, ok := value["host"]; ok && host != "" {
			rules[key].Host = host.(string)
		}

		if paths, ok := value["path"]; ok {
			rules[key].HTTP = &extensionsv1beta1.HTTPIngressRuleValue{
				Paths: path.Expand(paths.([]interface{})),
			}
		}
	}

	return rules
}

// Flatten structured object into unstructured.
func Flatten(in []extensionsv1beta1.IngressRule) []interface{} {
	flattened := make([]interface{}, len(in))

	for key, value := range in {
		row := map[string]interface{}{}

		if value.Host != "" {
			row["host"] = value.Host
		}

		if value.HTTP != nil {
			row["path"] = path.Flatten(value.HTTP.Paths)
		}

		flattened[key] = row
	}

	return flattened
}
