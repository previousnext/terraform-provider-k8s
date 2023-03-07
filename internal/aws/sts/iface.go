package sts

import (
	"context"

	signerv4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// ClientInterface provides an interface for the STS client.
type ClientInterface interface {
	GetCallerIdentity(ctx context.Context, params *sts.GetCallerIdentityInput, optFns ...func(options *sts.Options)) (*sts.GetCallerIdentityOutput, error)
}

// PresignClientInterface provides an interface for the STS Presign client.
type PresignClientInterface interface {
	PresignGetCallerIdentity(ctx context.Context, params *sts.GetCallerIdentityInput, optFns ...func(options *sts.PresignOptions)) (*signerv4.PresignedHTTPRequest, error)
}
