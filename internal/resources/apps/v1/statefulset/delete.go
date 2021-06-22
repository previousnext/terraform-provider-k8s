package statefulset

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Delete the Deployment.
func Delete(d *schema.ResourceData, m interface{}) error {
	conn := m.(*config.Client)

	namespace, name, err := id.Split(d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to get ID")
	}

	return conn.Kubernetes().AppsV1().StatefulSets(namespace).Delete(name, &metav1.DeleteOptions{})
}
