package clusterrole

import (
	"github.com/hashicorp/terraform/helper/schema"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/interfaceutils"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/rbac/v1/role/rule"
)

// Generate the ClusterRole.
func Generate(d *schema.ResourceData) (rbacv1.ClusterRole, error) {
	var (
		name      = d.Get(FieldName).(string)
		rawLabels = d.Get(FieldLabels).(map[string]interface{})
		rawRules  = d.Get(FieldRule).([]interface{})
	)

	role := rbacv1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Name:   name,
			Labels: interfaceutils.ExpandMap(rawLabels),
		},
		Rules: rule.Expand(rawRules),
	}

	return role, nil
}
