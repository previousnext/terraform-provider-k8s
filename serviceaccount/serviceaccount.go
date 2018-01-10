package serviceaccount

import (
	"github.com/hashicorp/terraform/helper/schema"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
				Description: "Name of the ServiceAccount.",
				Required:    true,
			},
			"namespace": {
				Type:        schema.TypeString,
				Description: "Namespace which the ServiceAccount resides",
				Required:    true,
			},
		},
	}
}

// Helper function for generating the ServiceAccount object.
func generateServiceAccount(d *schema.ResourceData) (corev1.ServiceAccount, error) {
	var (
		name      = d.Get("name").(string)
		namespace = d.Get("namespace").(string)
	)

	serviceaccount := corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
	}

	return serviceaccount, nil
}
