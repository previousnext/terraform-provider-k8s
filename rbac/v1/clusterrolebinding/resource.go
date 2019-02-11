package clusterrolebinding

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/previousnext/terraform-provider-k8s/apimachinery/objectmeta"
	"github.com/previousnext/terraform-provider-k8s/rbac/v1/rolebinding/subject"
)

const (
	// FieldRefKind identifies the ref_kind field.
	FieldRefKind = "ref_kind"
	// FieldRefName identifies the ref_kind field.
	FieldRefName = "ref_name"
	// FieldRefAPIGroup identifies the ref_kind field.
	FieldRefAPIGroup = "ref_api_group"
)

// Resource returns this packages Resource and Fields.
func Resource() *schema.Resource {
	return &schema.Resource{
		Create: Create,
		Read:   Read,
		Update: Update,
		Delete: Delete,

		Schema: map[string]*schema.Schema{
			objectmeta.FieldObjectMeta: objectmeta.Fields(),
			FieldRefKind: {
				Type:        schema.TypeString,
				Description: "Kind of Role being referenced",
				Required:    true,
			},
			FieldRefName: {
				Type:        schema.TypeString,
				Description: "Name of Role being referenced",
				Required:    true,
			},
			FieldRefAPIGroup: {
				Type:        schema.TypeString,
				Description: "API group of the Role being referenced",
				Optional:    true,
			},
			subject.FieldSubject: subject.Fields(),
		},
	}
}
