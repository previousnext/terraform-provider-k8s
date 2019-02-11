package rolebinding

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"

	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Create the RoleBinding.
func Create(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	binding, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	out, err := conn.RbacV1().RoleBindings(binding.ObjectMeta.Namespace).Create(&binding)
	if err != nil {
		return errors.Wrap(err, "failed to create")
	}

	d.SetId(id.Join(out.ObjectMeta))

	return nil
}
