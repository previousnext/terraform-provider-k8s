package resources

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/previousnext/terraform-provider-k8s/internal/resources/apiextensions/v1beta1/crd"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/apiregistration/v1beta1/apiservice"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/apps/v1/daemonset"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/apps/v1/deployment"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/apps/v1/statefulset"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/configmap"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/namespace"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/secret"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/service"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/serviceaccount"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/policy/v1beta1/poddisruptionbudget"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/rbac/v1/clusterrole"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/rbac/v1/clusterrolebinding"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/rbac/v1/role"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/rbac/v1/rolebinding"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/storage/v1/storageclass"
)

const (
	// FieldDeployment identifier for the Kubernetes Deployment.
	FieldDeployment = "k8s_apps_v1_deployment"
	// FieldDaemonSet identifier for the Kubernetes DaemonSet.
	FieldDaemonSet = "k8s_apps_v1_daemonset"
	// FieldStatefulSet identifier for the Kubernetes StatefulSet.
	FieldStatefulSet = "k8s_apps_v1_statefulset"

	// FieldStorageClass identifier for the Kubernetes StorageClass.
	FieldStorageClass = "k8s_storage_v1_storageclass"

	// FieldNamespace identifier for the Kubernetes Namespace.
	FieldNamespace = "k8s_core_v1_namespace"
	// FieldSecret identifier for the Kubernetes Secret.
	FieldSecret = "k8s_core_v1_secret"
	// FieldConfigMap identifier for the Kubernetes Secret.
	FieldConfigMap = "k8s_core_v1_configmap"
	// FieldServiceAccount identifier for the Kubernetes ServiceAccount.
	FieldServiceAccount = "k8s_core_v1_serviceaccount"
	// FieldService identifier for the Kubernetes Service.
	FieldService = "k8s_core_v1_service"

	// FieldRole identifier for the Kubernetes Role.
	FieldRole = "k8s_rbac_v1_role"
	// FieldRoleBinding identifier for the Kubernetes RoleBinding.
	FieldRoleBinding = "k8s_rbac_v1_rolebinding"
	// FieldClusterRole identifier for the Kubernetes ClusterRole.
	FieldClusterRole = "k8s_rbac_v1_clusterrole"
	// FieldClusterRoleBinding identifier for the Kubernetes ClusterRoleBinding.
	FieldClusterRoleBinding = "k8s_rbac_v1_clusterrolebinding"

	// FieldCustomtResourceDefinition identifier for the Kubernetes CustomtResourceDefinition.
	FieldCustomtResourceDefinition = "k8s_apiextensions_v1beta1_customresourcedefinition"

	// FieldAPIService identifier for the Kubernetes APIService.
	FieldAPIService = "k8s_apiregistration_v1_apiservice"

	// FieldPodDisruptionBudget identifier for the Kubernetes PodDisruptionBudget.
	FieldPodDisruptionBudget = "k8s_policy_v1beta1_poddisruptionbudget"
)

// Map returns a list of resources.
func Map() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		FieldNamespace:                 namespace.Resource(),
		FieldDeployment:                deployment.Resource(),
		FieldDaemonSet:                 daemonset.Resource(),
		FieldStatefulSet:               statefulset.Resource(),
		FieldStorageClass:              storageclass.Resource(),
		FieldServiceAccount:            serviceaccount.Resource(),
		FieldService:                   service.Resource(),
		FieldSecret:                    secret.Resource(),
		FieldConfigMap:                 configmap.Resource(),
		FieldRole:                      role.Resource(),
		FieldRoleBinding:               rolebinding.Resource(),
		FieldClusterRole:               clusterrole.Resource(),
		FieldClusterRoleBinding:        clusterrolebinding.Resource(),
		FieldCustomtResourceDefinition: crd.Resource(),
		FieldAPIService:                apiservice.Resource(),
		FieldPodDisruptionBudget:       poddisruptionbudget.Resource(),
	}
}
