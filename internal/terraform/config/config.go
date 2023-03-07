package config

import (
	"bytes"
	"context"
	"fmt"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/pkg/errors"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/previousnext/terraform-provider-k8s/internal/eks"
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
	ctx := context.TODO()

	cfg, err := awsconfig.LoadDefaultConfig(ctx,
		awsconfig.WithRegion(region),
		awsconfig.WithSharedConfigProfile(profile),
	)
	if err != nil {
		return "", fmt.Errorf("failed to get aws config: %w", err)
	}

	var (
		stsClient        = sts.NewFromConfig(cfg)
		stsPresignClient = sts.NewPresignClient(stsClient)
	)

	gen := eks.NewSTSTokenGenerator(stsPresignClient)
	if err != nil {
		return "", fmt.Errorf("failed to create token generator: %w", err)
	}

	token, err := gen.GenerateToken(ctx, cluster)
	if err != nil {
		return "", fmt.Errorf("failed to get sts token: %w", err)
	}

	return token, nil
}
