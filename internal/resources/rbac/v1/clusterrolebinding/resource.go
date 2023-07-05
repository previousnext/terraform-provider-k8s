package clusterrolebinding

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/previousnext/terraform-provider-k8s/internal/resources/rbac/v1/rolebinding/subject"
)

const (
	// FieldName is a field identifier.
	FieldName = "name"
	// FieldLabels is a field identifier.
	FieldLabels = "labels"
	// FieldRefKind is a field identifier.
	FieldRefKind = "ref_kind"
	// FieldRefName is a field identifier.
	FieldRefName = "ref_name"
	// FieldRefAPIGroup is a field identifier.
	FieldRefAPIGroup = "ref_api_group"
	// FieldSubject is a field identifier.
	FieldSubject = "subject"
)

// Resource returns this packages Resource and Fields.
func Resource() *schema.Resource {
	return &schema.Resource{
		CreateContext: Create,
		ReadContext:   Read,
		UpdateContext: Update,
		DeleteContext: Delete,

		Schema: map[string]*schema.Schema{
			FieldName: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			FieldLabels: &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			FieldRefKind: {
				Type:     schema.TypeString,
				Required: true,
			},
			FieldRefName: {
				Type:     schema.TypeString,
				Required: true,
			},
			FieldRefAPIGroup: {
				Type:     schema.TypeString,
				Optional: true,
			},
			FieldSubject: subject.Fields(),
		},
	}
}
