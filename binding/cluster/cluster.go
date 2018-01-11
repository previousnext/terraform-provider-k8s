package cluster

import (
	"github.com/hashicorp/terraform/helper/schema"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/binding/subject"
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
				Description: "Name of the RoleBinding.",
				Required:    true,
			},
			"ref_kind": {
				Type:        schema.TypeString,
				Description: "Kind of Role being referenced",
				Required:    true,
			},
			"ref_name": {
				Type:        schema.TypeString,
				Description: "Name of Role being referenced",
				Required:    true,
			},
			"ref_api_group": {
				Type:        schema.TypeString,
				Description: "API group of the Role being referenced",
				Optional:    true,
			},
			"subject": subject.Fields(),
		},
	}
}

// Helper function for generating the RoleBinding object.
func generateRoleBinding(d *schema.ResourceData) (rbacv1.ClusterRoleBinding, error) {
	var (
		name        = d.Get("name").(string)
		refKind     = d.Get("ref_kind").(string)
		refName     = d.Get("ref_name").(string)
		refAPIGroup = d.Get("ref_api_group").(string)
		subjects    = d.Get("subject").([]interface{})
	)

	role := rbacv1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		RoleRef: rbacv1.RoleRef{
			Kind:     refKind,
			Name:     refName,
			APIGroup: refAPIGroup,
		},
		Subjects: subject.Expand(subjects),
	}

	return role, nil
}
