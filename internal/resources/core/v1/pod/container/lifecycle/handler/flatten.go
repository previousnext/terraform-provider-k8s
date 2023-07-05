package handler

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/handler/exec"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/handler/http"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/handler/tcp"
)

// Flatten structured object into unstructured.
func Flatten(in *corev1.LifecycleHandler) []interface{} {
	out := make([]interface{}, 1)

	row := map[string]interface{}{}

	if in.HTTPGet != nil {
		row[FieldHTTP] = http.Flatten(in.HTTPGet)
	}

	if in.TCPSocket != nil {
		row[FieldTCP] = tcp.Flatten(in.TCPSocket)
	}

	if in.Exec != nil {
		row[FieldExec] = exec.Flatten(in.Exec)
	}

	out[0] = row

	return out
}
