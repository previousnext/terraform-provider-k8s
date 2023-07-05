package apiservice

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Read the APIService.
func Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	conn := m.(*config.Client)

	_, name, err := id.Split(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	service, err := conn.APIRegistration().ApiregistrationV1().APIServices().Get(ctx, name, metav1.GetOptions{})
	if kerrors.IsNotFound(err) {
		// This is how we tell Terraform that the resource does not exist.
		d.SetId("")
		return nil
	} else if err != nil {
		return diag.FromErr(err)
	}

	d.Set(FieldName, service.ObjectMeta.Name)
	d.Set(FieldLabels, service.ObjectMeta.Labels)

	d.Set(FieldGroup, service.Spec.Group)
	d.Set(FieldVersion, service.Spec.Version)

	d.Set(FieldServiceName, service.Spec.Service.Name)
	d.Set(FieldServiceNamespace, service.Spec.Service.Namespace)

	d.Set(FieldInsecureSkipTLSVerify, service.Spec.InsecureSkipTLSVerify)
	d.Set(FieldGroupPriorityMinimum, service.Spec.GroupPriorityMinimum)
	d.Set(FieldVersionPriority, service.Spec.VersionPriority)

	return diags
}
