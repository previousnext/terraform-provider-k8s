package container

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/args"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/command"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/envvar"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/lifecycle"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/mount"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/port"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/probe"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/resource"
)

// Expand will return a structured object.
func Expand(in []interface{}) ([]corev1.Container, error) {
	if len(in) == 0 {
		return []corev1.Container{}, nil
	}

	containers := make([]corev1.Container, len(in))

	for key, v := range in {
		value := v.(map[string]interface{})

		if name, ok := value[FieldName]; ok {
			containers[key].Name = name.(string)
		}

		if image, ok := value[FieldImage]; ok {
			containers[key].Image = image.(string)
		}

		if priv, ok := value[FieldPrivileged]; ok && priv == true {
			containers[key].SecurityContext = &corev1.SecurityContext{
				Privileged: pointerBool(priv.(bool)),
			}
		}

		if requests, ok := value[FieldRequests]; ok {
			expandedRequests, err := resource.Expand(requests.([]interface{}))
			if err != nil {
				return containers, err
			}

			containers[key].Resources.Requests = expandedRequests
		}

		if limits, ok := value[FieldLimits]; ok {
			expandedLimits, err := resource.Expand(limits.([]interface{}))
			if err != nil {
				return containers, err
			}

			containers[key].Resources.Limits = expandedLimits
		}

		if cmd, ok := value[FieldCommand]; ok {
			containers[key].Command = command.Expand(cmd.([]interface{}))
		}

		if cmd, ok := value[FieldArgs]; ok {
			containers[key].Args = args.Expand(cmd.([]interface{}))
		}

		if env, ok := value[FieldEnvVar]; ok {
			containers[key].Env = envvar.Expand(env.([]interface{}))
		}

		if prt, ok := value[FieldPort]; ok {
			ports, err := port.Expand(prt.([]interface{}))
			if err != nil {
				return containers, err
			}

			containers[key].Ports = ports
		}

		if liveness, ok := value[FieldLiveness]; ok {
			containers[key].LivenessProbe = probe.Expand(liveness.([]interface{}))
		}

		if readiness, ok := value[FieldReadiness]; ok {
			containers[key].ReadinessProbe = probe.Expand(readiness.([]interface{}))
		}

		if life, ok := value[FieldLifecycle]; ok {
			containers[key].Lifecycle = lifecycle.Expand(life.([]interface{}))
		}

		if mnt, ok := value[FieldMount]; ok {
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
