package namespace

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Delete the Namespace.
func Delete(d *schema.ResourceData, m interface{}) error {
	conn := m.(*config.Client)

	_, name, err := id.Split(d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to get ID")
	}

	return conn.Kubernetes().CoreV1().Namespaces().Delete(name, &metav1.DeleteOptions{})
}
