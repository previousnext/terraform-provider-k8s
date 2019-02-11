package clusterrole

import (
	"github.com/hashicorp/terraform/helper/schema"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/previousnext/terraform-provider-k8s/apimachinery/objectmeta"
	"github.com/previousnext/terraform-provider-k8s/rbac/v1/role/rule"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Read the ClusterRole.
func Read(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	_, name, err := id.Split(d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to get ID")
	}

	role, err := conn.RbacV1().ClusterRoles().Get(name, metav1.GetOptions{})
	if kerrors.IsNotFound(err) {
		// This is how we tell Terraform that the resource does not exist.
		d.SetId("")
		return nil
	} else if err != nil {
		return errors.Wrap(err, "failed to get")
	}

	d.Set(objectmeta.FieldObjectMeta, objectmeta.Flatten(role.ObjectMeta))
	d.Set(rule.FieldRule, rule.Flatten(role.Rules))

	return nil
}
