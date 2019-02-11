package ingress

import (
	"github.com/hashicorp/terraform/helper/schema"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"

	"github.com/previousnext/terraform-provider-k8s/apimachinery/objectmeta"
	"github.com/previousnext/terraform-provider-k8s/extensions/v1beta1/ingress/rule"
)

// Generate the Ingress.
func Generate(d *schema.ResourceData) (extensionsv1beta1.Ingress, error) {
	var (
		rawMeta = d.Get(objectmeta.FieldObjectMeta).([]interface{})
		rules   = d.Get(rule.FieldRule).([]interface{})
	)

	ingress := extensionsv1beta1.Ingress{
		ObjectMeta: objectmeta.Expand(rawMeta),
		Spec: extensionsv1beta1.IngressSpec{
			Rules: rule.Expand(rules),
		},
	}

	return ingress, nil
}
