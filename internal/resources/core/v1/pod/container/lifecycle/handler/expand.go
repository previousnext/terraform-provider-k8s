package handler

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/handler/exec"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/handler/http"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/handler/tcp"
)

// Expand will return a structured object.
func Expand(in []interface{}) *corev1.LifecycleHandler {
	if len(in) == 0 {
		return nil
	}

	handler := &corev1.LifecycleHandler{}

	raw := in[0].(map[string]interface{})

	if httpval, ok := raw[FieldHTTP]; ok {
		handler.HTTPGet = http.Expand(httpval.([]interface{}))
	}

	if tcpval, ok := raw[FieldTCP]; ok {
		handler.TCPSocket = tcp.Expand(tcpval.([]interface{}))
	}

	if execval, ok := raw[FieldExec]; ok {
		handler.Exec = exec.Expand(execval.([]interface{}))
	}

	return handler
}
