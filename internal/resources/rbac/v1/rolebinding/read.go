package rolebinding

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/resources/rbac/v1/rolebinding/subject"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Read the RoleBinding.
func Read(d *schema.ResourceData, m interface{}) error {
	conn := m.(*config.Client)

	namespace, name, err := id.Split(d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to get ID")
	}

	binding, err := conn.Kubernetes().RbacV1().RoleBindings(namespace).Get(name, metav1.GetOptions{})
	if kerrors.IsNotFound(err) {
		// This is how we tell Terraform that the resource does not exist.
		d.SetId("")
		return nil
	} else if err != nil {
		return errors.Wrap(err, "failed to get ID")
	}

	d.Set(FieldName, binding.ObjectMeta.Name)
	d.Set(FieldNamespace, binding.ObjectMeta.Namespace)
	d.Set(FieldLabels, binding.ObjectMeta.Labels)

	d.Set(FieldRefKind, binding.RoleRef.Kind)
	d.Set(FieldRefName, binding.RoleRef.Name)
	d.Set(FieldRefAPIGroup, binding.RoleRef.APIGroup)

	d.Set(FieldSubject, subject.Flatten(binding.Subjects))

	return nil
}
