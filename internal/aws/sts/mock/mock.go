package mock

import (
	"context"

	signerv4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sts"

	skprsts "github.com/previousnext/terraform-provider-k8s/internal/aws/sts"
)

// PresignClient is a mock client.
type PresignClient struct {
	skprsts.PresignClientInterface
}

// NewPresignClient creates a new mock client.
func NewPresignClient() *PresignClient {
	return &PresignClient{}
}

// PresignGetCallerIdentity implements the interface.
func (c *PresignClient) PresignGetCallerIdentity(ctx context.Context, params *sts.GetCallerIdentityInput, optFns ...func(options *sts.PresignOptions)) (*signerv4.PresignedHTTPRequest, error) {
	return &signerv4.PresignedHTTPRequest{
		URL: "http://example/com",
	}, nil
}
