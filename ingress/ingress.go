package ingress

import (
	"github.com/hashicorp/terraform/helper/schema"

	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/ingress/rule"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreate,
		Read:   resourceRead,
		Update: resourceUpdate,
		Delete: resourceDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the Ingress.",
				Required:    true,
			},
			"namespace": {
				Type:        schema.TypeString,
				Description: "Namespace which the Ingress will be run in.",
				Required:    true,
			},
			"rule": rule.Fields(),
		},
	}
}

func generateIngress(d *schema.ResourceData) (extensionsv1beta1.Ingress, error) {
	var (
		name      = d.Get("name").(string)
		namespace = d.Get("namespace").(string)
		rules     = d.Get("rule").([]interface{})
	)

	ingress := extensionsv1beta1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: extensionsv1beta1.IngressSpec{
			Rules: rule.Expand(rules),
		},
	}

	return ingress, nil
}
