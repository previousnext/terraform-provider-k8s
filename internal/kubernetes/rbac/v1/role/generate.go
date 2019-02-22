package role

import (
	"github.com/hashicorp/terraform/helper/schema"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/interfaceutils"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/rbac/v1/role/rule"
)

// Generate the Role.
func Generate(d *schema.ResourceData) (rbacv1.Role, error) {
	var (
		name      = d.Get(FieldName).(string)
		namespace = d.Get(FieldNamespace).(string)
		rawLabels = d.Get(FieldLabels).(map[string]interface{})
		rules     = d.Get(FieldRule).([]interface{})
	)

	role := rbacv1.Role{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    interfaceutils.ExpandMap(rawLabels),
		},
		Rules: rule.Expand(rules),
	}

	return role, nil
}
