package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mitchellh/go-homedir"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

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
			"k8s_cronjob": resourceCronJob(),
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
