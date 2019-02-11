package config

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Func which configures the Kubernetes provider.
// https://github.com/terraform-providers/terraform-provider-kubernetes/blob/master/kubernetes/provider.go
func Func(d *schema.ResourceData) (interface{}, error) {
	cfg := &rest.Config{}

	// Overriding with static configuration
	cfg.UserAgent = fmt.Sprintf("HashiCorp/1.0 Terraform/%s", terraform.VersionString())

	if v, ok := d.GetOk(FieldHost); ok {
		cfg.Host = v.(string)
	}

	if v, ok := d.GetOk(FieldInsecure); ok {
		cfg.Insecure = v.(bool)
	}

	if v, ok := d.GetOk(FieldClusterCACertificate); ok {
		cfg.CAData = bytes.NewBufferString(v.(string)).Bytes()
	}

	if v, ok := d.GetOk(FieldClientCertificate); ok {
		cfg.CertData = bytes.NewBufferString(v.(string)).Bytes()
	}

	if v, ok := d.GetOk(FieldClientKey); ok {
		cfg.KeyData = bytes.NewBufferString(v.(string)).Bytes()
	}

	k, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get Kubernetes config")
	}

	return k, nil
}
