package container

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/previousnext/terraform-provider-k8s/core/v1/pod/container/args"
	"github.com/previousnext/terraform-provider-k8s/core/v1/pod/container/command"
	"github.com/previousnext/terraform-provider-k8s/core/v1/pod/container/envvar"
	"github.com/previousnext/terraform-provider-k8s/core/v1/pod/container/mount"
	"github.com/previousnext/terraform-provider-k8s/core/v1/pod/container/resource"
)

// Expand will return a structured object.
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

func pointerBool(in bool) *bool {
	return &in
}