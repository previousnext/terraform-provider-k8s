package clusterrolebinding

import (
	"github.com/hashicorp/terraform/helper/schema"
	rbacv1 "k8s.io/api/rbac/v1"

	"github.com/previousnext/terraform-provider-k8s/apimachinery/objectmeta"
	"github.com/previousnext/terraform-provider-k8s/rbac/v1/rolebinding/subject"
)

// Generate the ClusterRoleBinding.
func Generate(d *schema.ResourceData) (rbacv1.ClusterRoleBinding, error) {
	var (
		rawMeta     = d.Get(objectmeta.FieldObjectMeta).([]interface{})
		refKind     = d.Get(FieldRefKind).(string)
		refName     = d.Get(FieldRefName).(string)
		refAPIGroup = d.Get(FieldRefAPIGroup).(string)
		rawSubjects = d.Get(subject.FieldSubject).([]interface{})
	)

	role := rbacv1.ClusterRoleBinding{
		ObjectMeta: objectmeta.Expand(rawMeta),
		RoleRef: rbacv1.RoleRef{
			Kind:     refKind,
			Name:     refName,
			APIGroup: refAPIGroup,
		},
		Subjects: subject.Expand(rawSubjects),
	}

	return role, nil
}
