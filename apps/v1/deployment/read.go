package deployment

import (
	"github.com/hashicorp/terraform/helper/schema"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/previousnext/terraform-provider-k8s/apimachinery/objectmeta"
	"github.com/previousnext/terraform-provider-k8s/core/v1/pod"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Read the Deployment.
func Read(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	namespace, name, err := id.Split(d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to get ID")
	}

	deployment, err := conn.AppsV1().Deployments(namespace).Get(name, metav1.GetOptions{})
	if kerrors.IsNotFound(err) {
		// This is how we tell Terraform that the resource does not exist.
		d.SetId("")
		return nil
	} else if err != nil {
		return errors.Wrap(err, "failed to get")
	}

	d.Set(objectmeta.FieldObjectMeta, objectmeta.Flatten(deployment.ObjectMeta))
	d.Set(FieldReplicas, *deployment.Spec.Replicas)
	d.Set(pod.FieldPod, pod.Flatten(deployment.Spec.Template.Spec))

	return nil
}
