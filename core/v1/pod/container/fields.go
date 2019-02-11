package container

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/previousnext/terraform-provider-k8s/core/v1/pod/container/args"
	"github.com/previousnext/terraform-provider-k8s/core/v1/pod/container/command"
	"github.com/previousnext/terraform-provider-k8s/core/v1/pod/container/envvar"
	"github.com/previousnext/terraform-provider-k8s/core/v1/pod/container/mount"
	"github.com/previousnext/terraform-provider-k8s/core/v1/pod/container/resource"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "List of containers to execute.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:        schema.TypeString,
					Description: "Name of this container.",
					Required:    true,
				},
				"image": { 
					Type:        schema.TypeString,
					Description: "Image to run for this container.",
					Required:    true,
				},
				"privileged": {
					Type:        schema.TypeBool, 
					Description: "Run this container as privileged.",
					Optional:    true,
				},
				"requests": resource.Fields(),
				"limits":   resource.Fields(),
				"command":  command.Fields(),
				"args":     args.Fields(),
				"env":      envvar.Fields(),
				"mount":    mount.Fields(),
			},
		},
	}
}
