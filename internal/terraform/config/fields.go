package config

import "github.com/hashicorp/terraform/helper/schema"

const (
	// FieldHost identifier for host field.
	FieldHost = "host"
	// FieldInsecure identifier for insecure field.
	FieldInsecure = "insecure"
	// FieldClusterCACertificate identifier for cluster_ca_certificate field.
	FieldClusterCACertificate = "cluster_ca_certificate"
	// FieldAWSRegion identifier for aws_region field.
	FieldAWSRegion = "aws_region"
	// FieldAWSProfile identifier for aws_profile field.
	FieldAWSProfile = "aws_profile"
	// FieldEKSCluster identifier for eks_cluster field.
	FieldEKSCluster = "eks_cluster"
)

// Fields which are used to configure the Kubernetes client.
func Fields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		FieldHost: {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("KUBE_HOST", ""),
			Description: "The hostname (in form of URI) of Kubernetes master.",
		},
		FieldInsecure: {
			Type:        schema.TypeBool,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("KUBE_INSECURE", false),
			Description: "Whether server should be accessed without verifying the TLS certificate.",
		},
		FieldEKSCluster: {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("AWS_PROFILE", ""),
			Description: "EKS cluster name.",
		},
		FieldAWSProfile: {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("AWS_PROFILE", ""),
			Description: "AWS Profile for authenticating with EKS.",
		},
		FieldAWSRegion: {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("AWS_PROFILE", ""),
			Description: "AWS region for connecting to EKS.",
		},
		FieldClusterCACertificate: {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("KUBE_CLUSTER_CA_CERT_DATA", ""),
			Description: "PEM-encoded root certificates bundle for TLS authentication.",
		},
	}
}
