package container

import (
	"github.com/previousnext/terraform-provider-k8s/core/v1/pod/container/envvar"
	"github.com/previousnext/terraform-provider-k8s/core/v1/pod/container/mount"
	"github.com/previousnext/terraform-provider-k8s/core/v1/pod/container/resource"
	corev1 "k8s.io/api/core/v1"
)

// Flatten structured object into unstructured.
func Flatten(in []corev1.Container) []interface{} {
	flattened := make([]interface{}, len(in))

	for key, value := range in {
		row := map[string]interface{}{}

		if value.Name != "" {
			row["name"] = value.Name
		}

		if value.Image != "" {
			row["image"] = value.Image
		}

		if value.SecurityContext != nil && value.SecurityContext.Privileged != nil {
			row["privileged"] = *value.SecurityContext.Privileged
		}

		if value.Resources.Requests != nil {
			row["requests"] = resource.Flatten(value.Resources.Requests)
		}

		if value.Resources.Limits != nil {
			row["limits"] = resource.Flatten(value.Resources.Limits)
		}

		if len(value.Command) > 0 {
			row["command"] = value.Command
		}

		if len(value.Args) > 0 {
			row["args"] = value.Args
		}

		if len(value.Env) > 0 {
			row["env"] = envvar.Flatten(value.Env)
		}

		if len(value.VolumeMounts) > 0 {
			row["mount"] = mount.Flatten(value.VolumeMounts)
		}

		flattened[key] = row
	}

	return flattened
}
