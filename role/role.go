package role

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/previousnext/terraform-provider-k8s/role/rule"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Resource returns this packages resource.
func Resource() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreate,
		Read:   resourceRead,
		Update: resourceUpdate,
		Delete: resourceDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the Role.",
				Required:    true,
			},
			"namespace": {
				Type:        schema.TypeString,
				Description: "Namespace which the Role resides",
				Required:    true,
			},
			"rule": rule.Fields(),
		},
	}
}

// Helper function for generating the Role object.
func generateRole(d *schema.ResourceData) (rbacv1.Role, error) {
	var (
		name      = d.Get("name").(string)
		namespace = d.Get("namespace").(string)
		rules     = d.Get("rule").([]interface{})
	)

	role := rbacv1.Role{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Rules: rule.Expand(rules),
	}

	return role, nil
}
