package handler

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/handler/exec"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/handler/http"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container/handler/tcp"
)

const (
	// FieldHTTP which is used to identify the HTTP field.
	FieldHTTP = "http"
	// FieldTCP which is used to identify the TCP field.
	FieldTCP = "tcp"
	// FieldExec which is used to identify the Exec field.
	FieldExec = "exec"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Description: "Event to occur during a containers lifecycle.",
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				FieldHTTP: http.Fields(),
				FieldTCP:  tcp.Fields(),
				FieldExec: exec.Fields(),
			},
		},
	}
}
