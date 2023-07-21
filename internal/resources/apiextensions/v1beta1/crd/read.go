package crd

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/resources/apiextensions/v1beta1/crd/names"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/apiextensions/v1beta1/crd/property"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Read the StorageClass.
func Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	conn := m.(*config.Client)

	_, name, err := id.Split(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	crd, err := conn.APIExtensions().ApiextensionsV1().CustomResourceDefinitions().Get(ctx, name, metav1.GetOptions{})
	if kerrors.IsNotFound(err) {
		// This is how we tell Terraform that the resource does not exist.
		d.SetId("")
		return nil
	} else if err != nil {
		return diag.FromErr(err)
	}

	d.Set(FieldName, crd.ObjectMeta.Name)
	d.Set(FieldLabels, crd.ObjectMeta.Labels)
	d.Set(FieldAnnotations, crd.ObjectMeta.Annotations)

	d.Set(FieldGroup, crd.Spec.Group)

	if len(crd.Spec.Versions) > 0 {
		d.Set(FieldVersion, crd.Spec.Versions[0].Name)
	}

	d.Set(FieldScope, crd.Spec.Scope)
	d.Set(FieldNames, names.Flatten(crd.Spec.Names))

	if len(crd.Spec.Versions) > 0 {
		d.Set(FieldProperty, property.Flatten(crd.Spec.Versions[0].Schema.OpenAPIV3Schema.Properties))
		d.Set(FieldRequired, crd.Spec.Versions[0].Schema.OpenAPIV3Schema.Required)
	}

	return diags
}
