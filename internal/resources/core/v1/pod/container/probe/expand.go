package probe

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/handler/exec"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/handler/http"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/handler/tcp"
)

// Expand will return a structured object.
func Expand(in []interface{}) *corev1.Probe {
	probe := &corev1.Probe{}

	if len(in) == 0 {
		return nil
	}

	raw := in[0].(map[string]interface{})

	if httpval, ok := raw[FieldHTTP]; ok {
		probe.HTTPGet = http.Expand(httpval.([]interface{}))
	}

	if tcpval, ok := raw[FieldTCP]; ok {
		probe.TCPSocket = tcp.Expand(tcpval.([]interface{}))
	}

	if execval, ok := raw[FieldExec]; ok {
		probe.Exec = exec.Expand(execval.([]interface{}))
	}

	if delay, ok := raw[FieldInitialDelaySeconds]; ok {
		probe.InitialDelaySeconds = int32(delay.(int))
	}

	if period, ok := raw[FieldPeriodSeconds]; ok {
		probe.PeriodSeconds = int32(period.(int))
	}

	if timeout, ok := raw[FieldTimeoutSeconds]; ok {
		probe.TimeoutSeconds = int32(timeout.(int))
	}

	if success, ok := raw[FieldSuccessThreshold]; ok {
		probe.SuccessThreshold = int32(success.(int))
	}

	if failure, ok := raw[FieldFailureThreshold]; ok {
		probe.FailureThreshold = int32(failure.(int))
	}

	return probe
}
