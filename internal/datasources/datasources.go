package datasources

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	deploymentimage "github.com/previousnext/terraform-provider-k8s/internal/datasources/apps/v1/deployment/image"
)

const (
	// FieldDeploymentImage identifier for the Kubernetes Deployment.
	FieldDeploymentImage = "k8s_apps_v1_deployment_image"
)

// Map returns a list of data sources.
func Map() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		FieldDeploymentImage: deploymentimage.Source(),
	}
}
