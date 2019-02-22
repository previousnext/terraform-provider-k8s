package config

import (
	"github.com/pkg/errors"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Client for interacting with Kubernetes resources.
type Client struct {
	kubernetes    *kubernetes.Clientset
	apiextensions *apiextensions.Clientset
}

// NewForConfig returns a Kubernetes client.
func NewForConfig(config *rest.Config) (*Client, error) {
	k, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get Kubernetes client")
	}

	a, err := apiextensions.NewForConfig(config)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get API Extensions client")
	}

	client := &Client{
		kubernetes:    k,
		apiextensions: a,
	}

	return client, nil
}

// Kubernetes ClientSet.
func (c *Client) Kubernetes() *kubernetes.Clientset {
	return c.kubernetes
}

// APIExtensions ClientSet.
func (c *Client) APIExtensions() *apiextensions.Clientset {
	return c.apiextensions
}
