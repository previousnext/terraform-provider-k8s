package container

import (
	"github.com/hashicorp/terraform/helper/schema"
	corev1 "k8s.io/api/core/v1"

	"github.com/previousnext/terraform-provider-k8s/container/args"
	"github.com/previousnext/terraform-provider-k8s/container/command"
	"github.com/previousnext/terraform-provider-k8s/container/envvar"
	"github.com/previousnext/terraform-provider-k8s/container/mount"
	"github.com/previousnext/terraform-provider-k8s/container/resource"
)

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

func Expand(in []interface{}) ([]corev1.Container, error) {
	if len(in) == 0 {
		return []corev1.Container{}, nil
	}

	containers := make([]corev1.Container, len(in))

	for key, v := range in {
		value := v.(map[string]interface{})

		if name, ok := value["name"]; ok {
			containers[key].Name = name.(string)
		}

		if image, ok := value["image"]; ok {
			containers[key].Image = image.(string)
		}

		if priv, ok := value["privileged"]; ok && priv == true {
			containers[key].SecurityContext = &corev1.SecurityContext{
				Privileged: pointerBool(priv.(bool)),
			}
		}

		if requests, ok := value["requests"]; ok {
			expandedRequests, err := resource.Expand(requests.([]interface{}))
			if err != nil {
				return containers, err
			}

			containers[key].Resources.Requests = expandedRequests
		}

		if limits, ok := value["limits"]; ok {
			expandedLimits, err := resource.Expand(limits.([]interface{}))
			if err != nil {
				return containers, err
			}

			containers[key].Resources.Limits = expandedLimits
		}

		if cmd, ok := value["command"]; ok {
			containers[key].Command = command.Expand(cmd.([]interface{}))
		}

		if cmd, ok := value["args"]; ok {
			containers[key].Args = args.Expand(cmd.([]interface{}))
		}

		if env, ok := value["env"]; ok {
			containers[key].Env = envvar.Expand(env.([]interface{}))
		}

		if mnt, ok := value["mount"]; ok {
			mounts, err := mount.Expand(mnt.([]interface{}))
			if err != nil {
				return containers, err
			}

			containers[key].VolumeMounts = mounts
		}
	}

	return containers, nil
}

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

func pointerBool(in bool) *bool {
	return &in
}
