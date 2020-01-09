package main

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/apiextensions/v1beta1/crd"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/apiregistration/v1beta1/apiservice"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/apps/v1/daemonset"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/apps/v1/deployment"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/apps/v1/statefulset"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/core/v1/configmap"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/core/v1/namespace"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/core/v1/secret"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/core/v1/service"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/core/v1/serviceaccount"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/policy/v1beta1/poddisruptionbudget"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/rbac/v1/clusterrole"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/rbac/v1/clusterrolebinding"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/rbac/v1/role"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/rbac/v1/rolebinding"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/storage/v1/storageclass"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
)

const (
	// ResourceDeployment identifier for the Kubernetes Deployment.
	ResourceDeployment = "k8s_apps_v1_deployment"
	// ResourceDaemonSet identifier for the Kubernetes DaemonSet.
	ResourceDaemonSet = "k8s_apps_v1_daemonset"
	// ResourceStatefulSet identifier for the Kubernetes StatefulSet.
	ResourceStatefulSet = "k8s_apps_v1_statefulset"

	// ResourceStorageClass identifier for the Kubernetes StorageClass.
	ResourceStorageClass = "k8s_storage_v1_storageclass"

	// ResourceNamespace identifier for the Kubernetes Namespace.
	ResourceNamespace = "k8s_core_v1_namespace"
	// ResourceSecret identifier for the Kubernetes Secret.
	ResourceSecret = "k8s_core_v1_secret"
	// ResourceConfigMap identifier for the Kubernetes Secret.
	ResourceConfigMap = "k8s_core_v1_configmap"
	// ResourceServiceAccount identifier for the Kubernetes ServiceAccount.
	ResourceServiceAccount = "k8s_core_v1_serviceaccount"
	// ResourceService identifier for the Kubernetes Service.
	ResourceService = "k8s_core_v1_service"

	// ResourceRole identifier for the Kubernetes Role.
	ResourceRole = "k8s_rbac_v1_role"
	// ResourceRoleBinding identifier for the Kubernetes RoleBinding.
	ResourceRoleBinding = "k8s_rbac_v1_rolebinding"
	// ResourceClusterRole identifier for the Kubernetes ClusterRole.
	ResourceClusterRole = "k8s_rbac_v1_clusterrole"
	// ResourceClusterRoleBinding identifier for the Kubernetes ClusterRoleBinding.
	ResourceClusterRoleBinding = "k8s_rbac_v1_clusterrolebinding"

	// ResourceCustomtResourceDefinition identifier for the Kubernetes CustomtResourceDefinition.
	ResourceCustomtResourceDefinition = "k8s_apiextensions_v1beta1_customresourcedefinition"

	// ResourceAPIService identifier for the Kubernetes APIService.
	ResourceAPIService = "k8s_apiregistration_v1beta1_apiservice"

	// ResourcePodDisruptionBudget identifier for the Kubernetes PodDisruptionBudget.
	ResourcePodDisruptionBudget = "k8s_policy_v1beta1_poddisruptionbudget"
)

// Provider returns this providers resources.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: config.Fields(),
		DataSourcesMap: map[string]*schema.Resource{
			ResourceService: service.Source(),
		},
		ResourcesMap: map[string]*schema.Resource{
			ResourceNamespace:                 namespace.Resource(),
			ResourceDeployment:                deployment.Resource(),
			ResourceDaemonSet:                 daemonset.Resource(),
			ResourceStatefulSet:               statefulset.Resource(),
			ResourceStorageClass:              storageclass.Resource(),
			ResourceServiceAccount:            serviceaccount.Resource(),
			ResourceService:                   service.Resource(),
			ResourceSecret:                    secret.Resource(),
			ResourceConfigMap:                 configmap.Resource(),
			ResourceRole:                      role.Resource(),
			ResourceRoleBinding:               rolebinding.Resource(),
			ResourceClusterRole:               clusterrole.Resource(),
			ResourceClusterRoleBinding:        clusterrolebinding.Resource(),
			ResourceCustomtResourceDefinition: crd.Resource(),
			ResourceAPIService:                apiservice.Resource(),
			ResourcePodDisruptionBudget:       poddisruptionbudget.Resource(),
		},
		ConfigureFunc: config.Func,
	}
}
