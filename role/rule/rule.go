package rule

import (
	"github.com/hashicorp/terraform/helper/schema"
	rbacv1 "k8s.io/api/rbac/v1"

	"github.com/previousnext/terraform-provider-k8s/role/rule/apigroups"
	"github.com/previousnext/terraform-provider-k8s/role/rule/resources"
	"github.com/previousnext/terraform-provider-k8s/role/rule/verbs"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Description: "Rules to apply to a Role",
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"api_groups": apigroups.Fields(),
				"resources":  resources.Fields(),
				"verbs":      verbs.Fields(),
			},
		},
	}
}

// Expand will return a structured object.
func Expand(in []interface{}) []rbacv1.PolicyRule {
	if len(in) == 0 {
		return []rbacv1.PolicyRule{}
	}

	rules := make([]rbacv1.PolicyRule, len(in))

	for key, v := range in {
		value := v.(map[string]interface{})

		if apiGroupsRaw, ok := value["api_groups"]; ok {
			rules[key].APIGroups = apigroups.Expand(apiGroupsRaw.([]interface{}))
		}

		if resourcesRaw, ok := value["resources"]; ok {
			rules[key].Resources = resources.Expand(resourcesRaw.([]interface{}))
		}

		if verbsRaw, ok := value["verbs"]; ok {
			rules[key].Verbs = verbs.Expand(verbsRaw.([]interface{}))
		}
	}

	return rules
}

// Flatten structured object into unstructured.
func Flatten(in []rbacv1.PolicyRule) []interface{} {
	flattened := make([]interface{}, len(in))

	for key, value := range in {
		row := map[string]interface{}{}

		if len(value.APIGroups) > 0 {
			row["api_groups"] = value.APIGroups
		}

		if len(value.Resources) > 0 {
			row["resources"] = value.Resources
		}

		if len(value.Verbs) > 0 {
			row["verbs"] = value.Verbs
		}

		flattened[key] = row
	}

	return flattened
}
