package crd

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
)

// Update the StorageClass.
func Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	conn := m.(*config.Client)

	c, err := Generate(d)
	if err != nil {
		return diag.FromErr(err)
	}

	crd, err := conn.APIExtensions().ApiextensionsV1().CustomResourceDefinitions().Get(ctx, c.ObjectMeta.Name, metav1.GetOptions{})
	if err != nil {
		return diag.FromErr(err)
	}

	crd.Spec = c.Spec

	_, err = conn.APIExtensions().ApiextensionsV1().CustomResourceDefinitions().Update(ctx, crd, metav1.UpdateOptions{})
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

