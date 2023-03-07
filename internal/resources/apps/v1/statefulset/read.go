package statefulset

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Read the Deployment.
func Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	conn := m.(*config.Client)

	namespace, name, err := id.Split(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	statefulset, err := conn.Kubernetes().AppsV1().StatefulSets(namespace).Get(ctx, name, metav1.GetOptions{})
	if kerrors.IsNotFound(err) {
		// This is how we tell Terraform that the resource does not exist.
		d.SetId("")
		return nil
	} else if err != nil {
		return diag.FromErr(err)
	}

	d.Set(FieldName, statefulset.ObjectMeta.Name)
	d.Set(FieldNamespace, statefulset.ObjectMeta.Namespace)
	d.Set(FieldLabels, statefulset.ObjectMeta.Labels)
	d.Set(FieldMatchLabels, statefulset.Spec.Selector.MatchLabels)
	d.Set(FieldReplicas, *statefulset.Spec.Replicas)

	d.Set(FieldPod, pod.Flatten(statefulset.Spec.Template))

	return diags
}
