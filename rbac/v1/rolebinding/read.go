package rolebinding

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/previousnext/terraform-provider-k8s/apimachinery/objectmeta"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
	"github.com/previousnext/terraform-provider-k8s/rbac/v1/rolebinding/subject"
)

// Read the RoleBinding.
func Read(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	namespace, name, err := id.Split(d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to get ID")
	}

	binding, err := conn.RbacV1().RoleBindings(namespace).Get(name, metav1.GetOptions{})
	if kerrors.IsNotFound(err) {
		// This is how we tell Terraform that the resource does not exist.
		d.SetId("")
		return nil
	} else if err != nil {
		return errors.Wrap(err, "failed to get ID")
	}

	d.Set(objectmeta.FieldObjectMeta, objectmeta.Flatten(binding.ObjectMeta))
	d.Set(FieldRefKind, binding.RoleRef.Kind)
	d.Set(FieldRefName, binding.RoleRef.Name)
	d.Set(FieldRefAPIGroup, binding.RoleRef.APIGroup)
	d.Set(subject.FieldSubject, subject.Flatten(binding.Subjects))

	return nil
}
