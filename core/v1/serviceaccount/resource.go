package serviceaccount

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/previousnext/terraform-provider-k8s/apimachinery/objectmeta"
)

// Resource returns this packages resource.
func Resource() *schema.Resource {
	return &schema.Resource{
		Create: Create,
		Read:   Read,
		Update: Update,
		Delete: Delete,

		Schema: map[string]*schema.Schema{
			objectmeta.FieldObjectMeta: objectmeta.Fields(),
		},
	}
}
