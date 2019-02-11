package serviceaccount

import (
	"github.com/hashicorp/terraform/helper/schema"
	corev1 "k8s.io/api/core/v1"

	"github.com/previousnext/terraform-provider-k8s/apimachinery/objectmeta"
)

// Generate the ServiceAccount.
func Generate(d *schema.ResourceData) (corev1.ServiceAccount, error) {
	var rawMeta = d.Get(objectmeta.FieldObjectMeta).([]interface{})

	serviceaccount := corev1.ServiceAccount{
		ObjectMeta: objectmeta.Expand(rawMeta),
	}

	return serviceaccount, nil
}
