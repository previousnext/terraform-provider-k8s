package main

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/previousnext/terraform-provider-k8s/apps/v1/daemonset"
	"github.com/previousnext/terraform-provider-k8s/apps/v1/deployment"
	"github.com/previousnext/terraform-provider-k8s/core/v1/serviceaccount"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
	"github.com/previousnext/terraform-provider-k8s/rbac/v1/clusterrole"
	"github.com/previousnext/terraform-provider-k8s/rbac/v1/clusterrolebinding"
	"github.com/previousnext/terraform-provider-k8s/rbac/v1/role"
	"github.com/previousnext/terraform-provider-k8s/rbac/v1/rolebinding"
	"github.com/previousnext/terraform-provider-k8s/storage/v1/storageclass"
)

const (
	// ResourceDeployment identifier for the Kubernetes Deployment.
	ResourceDeployment = "k8s_apps_v1_deployment"
	// ResourceDaemonSet identifier for the Kubernetes DaemonSet.
	ResourceDaemonSet = "k8s_apps_v1_daemonset"
	// ResourceStorageClass identifier for the Kubernetes StorageClass.
	ResourceStorageClass = "k8s_storage_v1_storageclass"
	// ResourceServiceAccount identifier for the Kubernetes ServiceAccount.
	ResourceServiceAccount = "k8s_core_v1_serviceaccount"
	// ResourceRole identifier for the Kubernetes Role.
	ResourceRole = "k8s_rbac_v1_role"
	// ResourceRoleBinding identifier for the Kubernetes RoleBinding.
	ResourceRoleBinding = "k8s_rbac_v1_rolebinding"
	// ResourceClusterRole identifier for the Kubernetes ClusterRole.
	ResourceClusterRole = "k8s_rbac_v1_clusterrole"
	// ResourceClusterRoleBinding identifier for the Kubernetes ClusterRoleBinding.
	ResourceClusterRoleBinding = "k8s_rbac_v1_clusterrolebinding"
)

// Provider returns this providers resources.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: config.Fields(),
		ResourcesMap: map[string]*schema.Resource{
			ResourceDeployment:         deployment.Resource(),
			ResourceDaemonSet:          daemonset.Resource(),
			ResourceStorageClass:       storageclass.Resource(),
			ResourceServiceAccount:     serviceaccount.Resource(),
			ResourceRole:               role.Resource(),
			ResourceRoleBinding:        rolebinding.Resource(),
			ResourceClusterRole:        clusterrole.Resource(),
			ResourceClusterRoleBinding: clusterrolebinding.Resource(),
		},
		ConfigureFunc: config.Func,
	}
}
