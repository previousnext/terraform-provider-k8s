package daemonset

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
)

// Update the DaemonSet.
func Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	conn := m.(*config.Client)

	daemonset, err := Generate(d)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = conn.Kubernetes().AppsV1().DaemonSets(daemonset.ObjectMeta.Namespace).Update(ctx, &daemonset, metav1.UpdateOptions{})
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
