package rule

import (
	rbacv1 "k8s.io/api/rbac/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/rbac/v1/role/rule/apigroups"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/rbac/v1/role/rule/resourcenames"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/rbac/v1/role/rule/resources"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/rbac/v1/role/rule/verbs"
)

// Expand will return a structured object.
func Expand(in []interface{}) []rbacv1.PolicyRule {
	if len(in) == 0 {
		return []rbacv1.PolicyRule{}
	}

	rules := make([]rbacv1.PolicyRule, len(in))

	for key, v := range in {
		value := v.(map[string]interface{})

		if apiGroupsRaw, ok := value[FieldAPIGroups]; ok {
			rules[key].APIGroups = apigroups.Expand(apiGroupsRaw.([]interface{}))
		}

		if resourcesRaw, ok := value[FieldResources]; ok {
			rules[key].Resources = resources.Expand(resourcesRaw.([]interface{}))
		}

		if resourceNamesRaw, ok := value[FieldResourceNames]; ok {
			rules[key].ResourceNames = resourcenames.Expand(resourceNamesRaw.([]interface{}))
		}

		if verbsRaw, ok := value[FieldVerbs]; ok {
			rules[key].Verbs = verbs.Expand(verbsRaw.([]interface{}))
		}
	}

	return rules
}
