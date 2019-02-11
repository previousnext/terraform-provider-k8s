package clusterrole

import (
	"github.com/hashicorp/terraform/helper/schema"
	rbacv1 "k8s.io/api/rbac/v1"

	"github.com/previousnext/terraform-provider-k8s/rbac/v1/role/rule"
	"github.com/previousnext/terraform-provider-k8s/apimachinery/objectmeta"
)

// Generate the ClusterRole.
func Generate(d *schema.ResourceData) (rbacv1.ClusterRole, error) {
	var (
		rawMeta  = d.Get(objectmeta.FieldObjectMeta).([]interface{}) 
		rawRules     = d.Get(rule.FieldRule).([]interface{})
	)

	role := rbacv1.ClusterRole{ 
		ObjectMeta: objectmeta.Expand(rawMeta),
		Rules: rule.Expand(rawRules),
	}

	return role, nil
}
