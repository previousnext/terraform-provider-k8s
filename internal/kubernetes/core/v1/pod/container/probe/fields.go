package probe

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/core/v1/pod/container/handler/exec"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/core/v1/pod/container/handler/http"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/core/v1/pod/container/handler/tcp"
)

const (
	// FieldHTTP which is used to identify the HTTP field.
	FieldHTTP = "http"
	// FieldTCP which is used to identify the TCP field.
	FieldTCP = "tcp"
	// FieldExec which is used to identify the Exec field.
	FieldExec = "exec"
	// FieldInitialDelaySeconds which is used to identify the initial delay field.
	FieldInitialDelaySeconds = "initial_delay_seconds"
	// FieldPeriodSeconds which is used to identify the period seconds field.
	FieldPeriodSeconds = "period_seconds"
	// FieldTimeoutSeconds which is used to identify the timeout seconds field.
	FieldTimeoutSeconds = "timeout_seconds"
	// FieldSuccessThreshold which is used to identify the success threshold field.
	FieldSuccessThreshold = "success_threshold"
	// FieldFailureThreshold which is used to identify the failure threshold field.
	FieldFailureThreshold = "failure_threshold"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Description: "Mount a volume into a path inside the container.",
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				FieldHTTP: http.Fields(),
				FieldTCP:  tcp.Fields(),
				FieldExec: exec.Fields(),
				FieldInitialDelaySeconds: {
					Description: "Number of seconds after the container has started before liveness or readiness probes are initiated.",
					Type:        schema.TypeInt,
					Optional:    true,
				},
				FieldPeriodSeconds: {
					Description: "How often (in seconds) to perform the probe.",
					Type:        schema.TypeInt,
					Optional:    true,
					Default:     10,
				},
				FieldTimeoutSeconds: {
					Description: "Number of seconds after which the probe times out.",
					Type:        schema.TypeInt,
					Optional:    true,
					Default:     1,
				},
				FieldSuccessThreshold: {
					Description: "Minimum consecutive successes for the probe to be considered successful after having failed.",
					Type:        schema.TypeInt,
					Optional:    true,
					Default:     1,
				},
				FieldFailureThreshold: {
					Description: "When a Pod starts and the probe fails, Kubernetes will try failureThreshold times before giving up. Giving up in case of liveness probe means restarting the Pod. In case of readiness probe the Pod will be marked Unready.",
					Type:        schema.TypeInt,
					Optional:    true,
					Default:     3,
				},
			},
		},
	}
}
