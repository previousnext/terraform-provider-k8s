package apiservice

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Create the APIService.
func Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	conn := m.(*config.Client)

	service, err := Generate(d)
	if err != nil {
		return diag.FromErr(err)
	}

	out, err := conn.APIRegistration().ApiregistrationV1().APIServices().Create(ctx, &service, metav1.CreateOptions{})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(id.Join(out.ObjectMeta))

	return diags
}
