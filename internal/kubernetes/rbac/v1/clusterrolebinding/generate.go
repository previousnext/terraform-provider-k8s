package clusterrolebinding

import (
	"github.com/hashicorp/terraform/helper/schema"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/interfaceutils"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/rbac/v1/rolebinding/subject"
)

// Generate the ClusterRoleBinding.
func Generate(d *schema.ResourceData) (rbacv1.ClusterRoleBinding, error) {
	var (
		name        = d.Get(FieldName).(string)
		rawLabels   = d.Get(FieldLabels).(map[string]interface{})
		refKind     = d.Get(FieldRefKind).(string)
		refName     = d.Get(FieldRefName).(string)
		refAPIGroup = d.Get(FieldRefAPIGroup).(string)
		rawSubjects = d.Get(FieldSubject).([]interface{})
	)

	role := rbacv1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:   name,
			Labels: interfaceutils.ExpandMap(rawLabels),
		},
		RoleRef: rbacv1.RoleRef{
			Kind:     refKind,
			Name:     refName,
			APIGroup: refAPIGroup,
		},
		Subjects: subject.Expand(rawSubjects),
	}

	return role, nil
}
