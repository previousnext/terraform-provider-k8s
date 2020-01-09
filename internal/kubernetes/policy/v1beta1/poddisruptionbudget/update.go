package poddisruptionbudget

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"

	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
)

// Update the PodDisruptionBudget.
func Update(d *schema.ResourceData, m interface{}) error {
	conn := m.(*config.Client)

	budget, err := Generate(d)
	if err != nil {
		return errors.Wrap(err, "failed to generate")
	}

	_, err = conn.Kubernetes().PolicyV1beta1().PodDisruptionBudgets(budget.ObjectMeta.Namespace).Update(&budget)
	if err != nil {
		return errors.Wrap(err, "failed to update")
	}

	return nil
}
