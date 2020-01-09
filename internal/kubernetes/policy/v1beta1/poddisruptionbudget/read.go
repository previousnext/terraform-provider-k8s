package poddisruptionbudget

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Read the PodDisruptionBudget.
func Read(d *schema.ResourceData, m interface{}) error {
	conn := m.(*config.Client)

	namespace, name, err := id.Split(d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to get ID")
	}

	budget, err := conn.Kubernetes().PolicyV1beta1().PodDisruptionBudgets(namespace).Get(name, metav1.GetOptions{})
	if kerrors.IsNotFound(err) {
		// This is how we tell Terraform that the resource does not exist.
		d.SetId("")
		return nil
	} else if err != nil {
		return errors.Wrap(err, "failed to get")
	}

	d.Set(FieldName, budget.ObjectMeta.Name)
	d.Set(FieldNamespace, budget.ObjectMeta.Namespace)
	d.Set(FieldLabels, budget.ObjectMeta.Labels)
	d.Set(FieldMinAvailable, budget.Spec.MinAvailable.String())
	d.Set(FieldMatchLabels, budget.Spec.Selector.MatchLabels)

	return nil
}
