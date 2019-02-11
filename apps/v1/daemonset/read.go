package daemonset

import (
	"github.com/hashicorp/terraform/helper/schema"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"github.com/pkg/errors"

	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
	"github.com/previousnext/terraform-provider-k8s/apimachinery/objectmeta"
	"github.com/previousnext/terraform-provider-k8s/core/v1/pod"
)

// Read the DaemonSet.
func Read(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	namespace, name, err := id.Split(d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to get ID")
	}

	daemonset, err := conn.AppsV1().DaemonSets(namespace).Get(name, metav1.GetOptions{})
	if kerrors.IsNotFound(err) {
		// This is how we tell Terraform that the resource does not exist.
		d.SetId("")
		return nil
	} else if err != nil {
		return errors.Wrap(err, "failed to get")
	}

	d.Set(objectmeta.FieldObjectMeta, objectmeta.Flatten(daemonset.ObjectMeta))
	d.Set(pod.FieldPod, pod.Flatten(daemonset.Spec.Template.Spec))

	return nil
}
