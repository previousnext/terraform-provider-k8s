package config

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/kubernetes-sigs/aws-iam-authenticator/pkg/token"
	"github.com/pkg/errors"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Func which configures the Kubernetes provider.
// https://github.com/terraform-providers/terraform-provider-kubernetes/blob/master/kubernetes/provider.go
func Func(d *schema.ResourceData) (interface{}, error) {
	if v, ok := d.GetOk(FieldKubeConfig); ok {
		cfg, err := clientcmd.BuildConfigFromFlags("", v.(string))
		if err != nil {
			return nil, err
		}

		return NewForConfig(cfg)
	}

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

	if v, ok := d.GetOk(FieldEKSCluster); ok {
		var (
			region  string
			profile string
			cluster = v.(string)
		)

		if v, ok := d.GetOk(FieldAWSProfile); ok {
			profile = v.(string)
		}

		if v, ok := d.GetOk(FieldAWSRegion); ok {
			region = v.(string)
		}

		token, err := eksToken(region, cluster, profile)
		if err != nil {
			return nil, errors.Wrap(err, "failed to EKS bearer token")
		}

		cfg.BearerToken = token
	}

	k, err := NewForConfig(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get client")
	}

	return k, nil
}

// Helper function to generate an EKS Kubernetes client.
func eksToken(region, cluster, profile string) (string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewSharedCredentials("", profile),
	})
	if err != nil {
		return "", errors.Wrap(err, "failed to get session")
	}

	gen, err := token.NewGenerator(false)
	if err != nil {
		return "", errors.Wrap(err, "failed to create token client")
	}

	token, err := gen.GetWithSTS(cluster, sts.New(sess))
	if err != nil {
		return "", errors.Wrap(err, "failed to get token")
	}

	return token.Token, nil
}
