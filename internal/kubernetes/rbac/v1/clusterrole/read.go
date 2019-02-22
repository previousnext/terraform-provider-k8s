package clusterrole

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/rbac/v1/role/rule"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Read the ClusterRole.
func Read(d *schema.ResourceData, m interface{}) error {
	conn := m.(*config.Client)

	_, name, err := id.Split(d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to get ID")
	}

	role, err := conn.Kubernetes().RbacV1().ClusterRoles().Get(name, metav1.GetOptions{})
	if kerrors.IsNotFound(err) {
		// This is how we tell Terraform that the resource does not exist.
		d.SetId("")
		return nil
	} else if err != nil {
		return errors.Wrap(err, "failed to get")
	}

	d.Set(FieldName, role.ObjectMeta.Name)
	d.Set(FieldLabels, role.ObjectMeta.Labels)

	d.Set(FieldRule, rule.Flatten(role.Rules))

	return nil
}
