package datasources

import (
	"github.com/hashicorp/terraform/helper/schema"

	deploymentimage "github.com/previousnext/terraform-provider-k8s/internal/datasources/apps/v1/deployment/image"
)

const (
	// FieldDeploymentImage identifier for the Kubernetes Deployment.
	FieldDeploymentImage = "k8s_apps_v1_deployment_image"
)

// DataSourcesMap returns a list of data sources.
func DataSourcesMap() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		FieldDeploymentImage: deploymentimage.Source(),
	}
}
