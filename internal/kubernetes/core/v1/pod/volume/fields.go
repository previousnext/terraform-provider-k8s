package volume

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/core/v1/pod/volume/nfs"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/core/v1/pod/volume/hostpath"
)

const (
	// FieldName is a field identifier.
	FieldName = "name"
	// FieldPVC is a field identifier.
	FieldPVC = "pvc"
	// FieldConfigMap is a field identifier.
	FieldConfigMap = "configmap"
	// FieldEmptyDir is a field identifier.
	FieldEmptyDir = "empty_dir"
	// FieldNFS is a field identifier.
	FieldNFS = "nfs"
	// FieldHostPath is a field identifier.
	FieldHostPath = "host_path"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "List of Volumes used to mount into containers.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				FieldName: &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				FieldPVC: {
					Type:     schema.TypeString,
					Optional: true,
				},
				FieldConfigMap: {
					Type:     schema.TypeString,
					Optional: true,
				},
				FieldEmptyDir: {
					Type:     schema.TypeString,
					Optional: true,
				},
				FieldNFS:      nfs.Fields(),
				FieldHostPath: hostpath.Fields(),
			},
		},
	}
}
