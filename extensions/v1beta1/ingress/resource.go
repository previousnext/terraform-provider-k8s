package ingress

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/previousnext/terraform-provider-k8s/apimachinery/objectmeta"
	"github.com/previousnext/terraform-provider-k8s/extensions/v1beta1/ingress/rule"
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
			rule.FieldRule:             rule.Fields(),
		},
	}
}
