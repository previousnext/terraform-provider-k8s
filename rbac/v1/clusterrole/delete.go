package clusterrole

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Delete the ClusterRole.
func Delete(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	_, name, err := id.Split(d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to delete")
	}

	return conn.RbacV1().ClusterRoles().Delete(name, &metav1.DeleteOptions{})
}
