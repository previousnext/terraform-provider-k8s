package cluster

import (
	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/previousnext/terraform-provider-k8s/binding/subject"
	"github.com/previousnext/terraform-provider-k8s/utils/id"
)

func resourceRead(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	_, name, err := id.Split(d.Id())
	if err != nil {
		return err
	}

	binding, err := conn.RbacV1().ClusterRoleBindings().Get(name, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		// This is how we tell Terraform that the resource does not exist.
		d.SetId("")
		return nil
	} else if err != nil {
		return err
	}

	d.Set("name", binding.ObjectMeta.Name)
	d.Set("ref_kind", binding.RoleRef.Kind)
	d.Set("ref_name", binding.RoleRef.Name)
	d.Set("ref_api_group", binding.RoleRef.APIGroup)
	d.Set("subject", subject.Flatten(binding.Subjects))

	return nil
}
