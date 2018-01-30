package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mitchellh/go-homedir"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	clusterrolebinding "github.com/previousnext/terraform-provider-k8s/binding/cluster"
	rolebinding "github.com/previousnext/terraform-provider-k8s/binding/role"
	"github.com/previousnext/terraform-provider-k8s/cronjob"
	"github.com/previousnext/terraform-provider-k8s/daemonset"
	"github.com/previousnext/terraform-provider-k8s/deployment"
	"github.com/previousnext/terraform-provider-k8s/ingress"
	"github.com/previousnext/terraform-provider-k8s/role"
	clusterrole "github.com/previousnext/terraform-provider-k8s/role/cluster"
	"github.com/previousnext/terraform-provider-k8s/serviceaccount"
	"github.com/previousnext/terraform-provider-k8s/storageclass"
)

// Provider returns this providers resources.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"context": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Context to use for authentication",
			},
			"kubeconfig": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "~/.kube/config",
				Description: "Path to the Kubernetes config file",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"k8s_cronjob":              cronjob.Resource(),
			"k8s_deployment":           deployment.Resource(),
			"k8s_daemonset":            daemonset.Resource(),
			"k8s_storageclass":         storageclass.Resource(),
			"k8s_ingress":              ingress.Resource(),
			"k8s_service_account":      serviceaccount.Resource(),
			"k8s_role":                 role.Resource(),
			"k8s_role_binding":         rolebinding.Resource(),
			"k8s_cluster_role":         clusterrole.Resource(),
			"k8s_cluster_role_binding": clusterrolebinding.Resource(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	var (
		context    = d.Get("context").(string)
		kubeconfig = d.Get("kubeconfig").(string)
	)

	path, err := homedir.Expand(kubeconfig)
	if err != nil {
		return nil, err
	}

	loader := &clientcmd.ClientConfigLoadingRules{
		ExplicitPath: path,
	}

	overrides := &clientcmd.ConfigOverrides{
		CurrentContext: context,
	}

	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loader, overrides).ClientConfig()
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(config)
}
