package role

import (
	"github.com/hashicorp/terraform/helper/schema"
	rbacv1 "k8s.io/api/rbac/v1"

	"github.com/previousnext/terraform-provider-k8s/apimachinery/objectmeta"
	"github.com/previousnext/terraform-provider-k8s/rbac/v1/role/rule"
)

// Generate the Role.
func Generate(d *schema.ResourceData) (rbacv1.Role, error) {
	var (
		rawMeta = d.Get(objectmeta.FieldObjectMeta).([]interface{})
		rules   = d.Get(rule.FieldRule).([]interface{})
	)

	role := rbacv1.Role{
		ObjectMeta: objectmeta.Expand(rawMeta),
		Rules:      rule.Expand(rules),
	}

	return role, nil
}
