package daemonset

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/core/v1/pod"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Read the DaemonSet.
func Read(d *schema.ResourceData, m interface{}) error {
	conn := m.(*config.Client)

	namespace, name, err := id.Split(d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to get ID")
	}

	daemonset, err := conn.Kubernetes().AppsV1().DaemonSets(namespace).Get(name, metav1.GetOptions{})
	if kerrors.IsNotFound(err) {
		// This is how we tell Terraform that the resource does not exist.
		d.SetId("")
		return nil
	} else if err != nil {
		return errors.Wrap(err, "failed to get")
	}

	d.Set(FieldName, daemonset.ObjectMeta.Name)
	d.Set(FieldNamespace, daemonset.ObjectMeta.Namespace)
	d.Set(FieldLabels, daemonset.ObjectMeta.Labels)
	d.Set(FieldMatchLabels, daemonset.Spec.Selector.MatchLabels)
	d.Set(FieldPod, pod.Flatten(daemonset.Spec.Template))

	return nil
}
