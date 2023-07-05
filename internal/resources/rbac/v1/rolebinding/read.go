package rolebinding

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/resources/rbac/v1/rolebinding/subject"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Read the RoleBinding.
func Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	conn := m.(*config.Client)

	namespace, name, err := id.Split(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	binding, err := conn.Kubernetes().RbacV1().RoleBindings(namespace).Get(ctx, name, metav1.GetOptions{})
	if kerrors.IsNotFound(err) {
		// This is how we tell Terraform that the resource does not exist.
		d.SetId("")
		return nil
	} else if err != nil {
		return diag.FromErr(err)
	}

	d.Set(FieldName, binding.ObjectMeta.Name)
	d.Set(FieldNamespace, binding.ObjectMeta.Namespace)
	d.Set(FieldLabels, binding.ObjectMeta.Labels)

	d.Set(FieldRefKind, binding.RoleRef.Kind)
	d.Set(FieldRefName, binding.RoleRef.Name)
	d.Set(FieldRefAPIGroup, binding.RoleRef.APIGroup)

	d.Set(FieldSubject, subject.Flatten(binding.Subjects))

	return diags
}
