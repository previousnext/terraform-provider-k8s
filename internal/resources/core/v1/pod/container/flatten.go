package container

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/envvar"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/lifecycle"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/mount"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/port"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/probe"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/resource"
)

// Flatten structured object into unstructured.
func Flatten(in []corev1.Container) []interface{} {
	flattened := make([]interface{}, len(in))

	for key, value := range in {
		row := map[string]interface{}{}

		if value.Name != "" {
			row[FieldName] = value.Name
		}

		if value.Image != "" {
			row[FieldImage] = value.Image
		}

		if value.SecurityContext != nil && value.SecurityContext.Privileged != nil {
			row[FieldPrivileged] = *value.SecurityContext.Privileged
		}

		if value.Resources.Requests != nil {
			row[FieldRequests] = resource.Flatten(value.Resources.Requests)
		}

		if value.Resources.Limits != nil {
			row[FieldLimits] = resource.Flatten(value.Resources.Limits)
		}

		if len(value.Command) > 0 {
			row[FieldCommand] = value.Command
		}

		if len(value.Args) > 0 {
			row[FieldArgs] = value.Args
		}

		if len(value.Env) > 0 {
			row[FieldEnvVar] = envvar.Flatten(value.Env)
		}

		if len(value.Ports) > 0 {
			row[FieldPort] = port.Flatten(value.Ports)
		}

		if len(value.VolumeMounts) > 0 {
			row[FieldMount] = mount.Flatten(value.VolumeMounts)
		}

		if value.LivenessProbe != nil {
			row[FieldLiveness] = probe.Flatten(value.LivenessProbe)
		}

		if value.ReadinessProbe != nil {
			row[FieldReadiness] = probe.Flatten(value.ReadinessProbe)
		}

		if value.Lifecycle != nil {
			row[FieldLifecycle] = lifecycle.Flatten(value.Lifecycle)
		}

		flattened[key] = row
	}

	return flattened
}
