package apiservice

import (
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	// FieldName is a field identifier.
	FieldName = "name"
	// FieldLabels is a field identifier.
	FieldLabels = "labels"
	// FieldGroup is a field identifier.
	FieldGroup = "group"
	// FieldVersion is a field identifier.
	FieldVersion = "version"
	// FieldServiceName is a field identifier.
	FieldServiceName = "service_name"
	// FieldServiceNamespace is a field identifier.
	FieldServiceNamespace = "service_namespace"
	// FieldInsecureSkipTLSVerify is a field identifier.
	FieldInsecureSkipTLSVerify = "insecure_skip_tls_verify"
	// FieldGroupPriorityMinimum is a field identifier.
	FieldGroupPriorityMinimum = "group_priority_minimum"
	// FieldVersionPriority is a field identifier.
	FieldVersionPriority = "version_priority"
)

// Resource returns this packages Resource and Fields.
func Resource() *schema.Resource {
	return &schema.Resource{
		Create: Create,
		Read:   Read,
		Update: Update,
		Delete: Delete,

		Schema: map[string]*schema.Schema{
			FieldName: {
				Type:     schema.TypeString,
				Required: true,
			},
			FieldLabels: {
				Type:     schema.TypeMap,
				Optional: true,
			},
			FieldServiceName: {
				Type:     schema.TypeString,
				Required: true,
			},
			FieldServiceNamespace: {
				Type:     schema.TypeString,
				Required: true,
			},
			FieldGroup: {
				Type:     schema.TypeString,
				Required: true,
			},
			FieldVersion: {
				Type:     schema.TypeString,
				Required: true,
			},
			FieldInsecureSkipTLSVerify: {
				Type:     schema.TypeBool,
				Required: true,
			},
			FieldGroupPriorityMinimum: {
				Type:     schema.TypeInt,
				Required: true,
			},
			FieldVersionPriority: {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}
