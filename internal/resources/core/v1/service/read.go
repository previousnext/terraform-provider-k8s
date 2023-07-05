package service

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/service/ingress"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/service/port"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Read the ServiceAccount.
func Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	conn := m.(*config.Client)

	namespace, name, err := id.Split(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	service, err := conn.Kubernetes().CoreV1().Services(namespace).Get(ctx, name, metav1.GetOptions{})
	if kerrors.IsNotFound(err) {
		// This is how we tell Terraform that the resource does not exist.
		d.SetId("")
		return nil
	} else if err != nil {
		return diag.FromErr(err)
	}

	d.Set(FieldName, service.ObjectMeta.Name)
	d.Set(FieldNamespace, service.ObjectMeta.Namespace)
	d.Set(FieldLabels, service.ObjectMeta.Labels)
	d.Set(FieldAnnotations, service.ObjectMeta.Annotations)
	d.Set(FieldType, service.Spec.Type)
	d.Set(FieldPort, port.Flatten(service.Spec.Ports))
	d.Set(FieldSelector, service.Spec.Selector)
	d.Set(FieldIngress, ingress.Flatten(service.Status.LoadBalancer.Ingress))

	return diags
}
