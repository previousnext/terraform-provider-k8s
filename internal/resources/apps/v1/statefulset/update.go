package statefulset

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
)

// Update the Deployment.
func Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	conn := m.(*config.Client)

	statefulset, err := Generate(d)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = conn.Kubernetes().AppsV1().StatefulSets(statefulset.ObjectMeta.Namespace).Update(ctx, &statefulset, metav1.UpdateOptions{})
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
