package image

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Read the Deployment.
func Read(d *schema.ResourceData, m interface{}) error {
	conn := m.(*config.Client)

	var (
		namespace = d.Get(FieldNamespace).(string)
		name      = d.Get(FieldName).(string)
		container = d.Get(FieldContainer).(string)
		fallback  = d.Get(FieldFallback).(string)
	)

	d.SetId(id.Join(metav1.ObjectMeta{
		Namespace: namespace,
		Name:      name,
	}))

	deployment, err := conn.Kubernetes().AppsV1().Deployments(namespace).Get(name, metav1.GetOptions{})
	if kerrors.IsNotFound(err) {
		d.Set(FieldResult, fallback)
		return nil
	} else if err != nil {
		return errors.Wrap(err, "failed to get")
	}

	getImage := func(containers []corev1.Container, name, fallback string) string {
		for _, c := range deployment.Spec.Template.Spec.Containers {
			if c.Name == container {
				return c.Image
			}
		}

		return fallback
	}

	result := getImage(deployment.Spec.Template.Spec.Containers, container, fallback)

	d.Set(FieldResult, result)

	return nil
}
