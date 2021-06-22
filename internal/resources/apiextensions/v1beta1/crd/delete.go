package crd

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Delete the StorageClass.
func Delete(d *schema.ResourceData, m interface{}) error {
	// We don't delete here. This generally results in very destructive outcomes.
	return nil
}
