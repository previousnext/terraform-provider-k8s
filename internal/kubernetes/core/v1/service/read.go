package service

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/core/v1/service/ingress"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/core/v1/service/port"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Read the ServiceAccount.
func Read(d *schema.ResourceData, m interface{}) error {
	conn := m.(*config.Client)

	namespace, name, err := id.Split(d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to get ID")
	}

	service, err := conn.Kubernetes().CoreV1().Services(namespace).Get(name, metav1.GetOptions{})
	if kerrors.IsNotFound(err) {
		// This is how we tell Terraform that the resource does not exist.
		d.SetId("")
		return nil
	} else if err != nil {
		return errors.Wrap(err, "failed to get")
	}

	d.Set(FieldName, service.ObjectMeta.Name)
	d.Set(FieldNamespace, service.ObjectMeta.Namespace)
	d.Set(FieldLabels, service.ObjectMeta.Labels)
	d.Set(FieldType, service.Spec.Type)
	d.Set(FieldPort, port.Flatten(service.Spec.Ports))
	d.Set(FieldSelector, service.Spec.Selector)
	d.Set(FieldIngress, ingress.Flatten(service.Status.LoadBalancer.Ingress))

	return nil
}
