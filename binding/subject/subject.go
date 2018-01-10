package subject

import (
	"github.com/hashicorp/terraform/helper/schema"
	rbacv1 "k8s.io/api/rbac/v1"
)

func Fields() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Subjects which will receive this RoleBinding.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"kind": {
					Type:        schema.TypeString,
					Description: "Kind eg. User, Group, ServiceAccount etc",
					Required:    true,
				},
				"name": {
					Type:        schema.TypeString,
					Description: "Name of the allowed",
					Required:    true,
				},
				"api_group": {
					Type:        schema.TypeString,
					Description: "API group of the allowed",
					Optional:    true,
				},
				"namespace": {
					Type:        schema.TypeString,
					Description: "Namespace of the allowed",
					Optional:    true,
				},
			},
		},
	}
}

func Expand(in []interface{}) []rbacv1.Subject {
	if len(in) == 0 {
		return []rbacv1.Subject{}
	}

	rules := make([]rbacv1.Subject, len(in))

	for key, v := range in {
		value := v.(map[string]interface{})

		if kind, ok := value["kind"]; ok {
			rules[key].Kind = kind.(string)
		}

		if name, ok := value["name"]; ok {
			rules[key].Name = name.(string)
		}

		if apiGroup, ok := value["api_group"]; ok {
			rules[key].APIGroup = apiGroup.(string)
		}

		if namespace, ok := value["namespace"]; ok {
			rules[key].Namespace = namespace.(string)
		}
	}

	return rules
}

func Flatten(in []rbacv1.Subject) []interface{} {
	flattened := make([]interface{}, len(in))

	for key, value := range in {
		row := map[string]interface{}{}

		if value.Kind != "" {
			row["kind"] = value.Kind
		}

		if value.Name != "" {
			row["name"] = value.Name
		}

		if value.APIGroup != "" {
			row["api_group"] = value.APIGroup
		}

		if value.Namespace != "" {
			row["namespace"] = value.Namespace
		}

		flattened[key] = row
	}

	return flattened
}
