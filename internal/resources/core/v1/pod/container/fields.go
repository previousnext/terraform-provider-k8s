package container

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/args"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/command"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/envvar"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/lifecycle"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/mount"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/port"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/probe"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/resource"
)

const (
	// FieldName is a field identifier.
	FieldName = "name"
	// FieldImage is a field identifier.
	FieldImage = "image"
	// FieldPrivileged is a field identifier.
	FieldPrivileged = "privileged"
	// FieldRequests is a field identifier.
	FieldRequests = "requests"
	// FieldLimits is a field identifier.
	FieldLimits = "limits"
	// FieldCommand is a field identifier.
	FieldCommand = "command"
	// FieldArgs is a field identifier.
	FieldArgs = "args"
	// FieldEnvVar is a field identifier.
	FieldEnvVar = "envvar"
	// FieldPort is a field identifier.
	FieldPort = "port"
	// FieldReadiness is a field identifier.
	FieldReadiness = "readiness_probe"
	// FieldLiveness is a field identifier.
	FieldLiveness = "liveness_probe"
	// FieldLifecycle is a field identifier.
	FieldLifecycle = "lifecycle"
	// FieldMount is a field identifier.
	FieldMount = "mount"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Description: "List of containers to execute.",
		Type:        schema.TypeList,
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				FieldName: {
					Type:        schema.TypeString,
					Description: "Name of this container.",
					Required:    true,
				},
				FieldImage: {
					Type:        schema.TypeString,
					Description: "Image to run for this container.",
					Required:    true,
				},
				FieldPrivileged: {
					Type:        schema.TypeBool,
					Description: "Run this container as privileged.",
					Optional:    true,
				},
				FieldRequests:  resource.Fields(),
				FieldLimits:    resource.Fields(),
				FieldCommand:   command.Fields(),
				FieldArgs:      args.Fields(),
				FieldEnvVar:    envvar.Fields(),
				FieldPort:      port.Fields(),
				FieldReadiness: probe.Fields(),
				FieldLiveness:  probe.Fields(),
				FieldLifecycle: lifecycle.Fields(),
				FieldMount:     mount.Fields(),
			},
		},
	}
}
