package eks

import (
	"context"
	"encoding/base64"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	smithyhttp "github.com/aws/smithy-go/transport/http"

	awssts "github.com/previousnext/terraform-provider-k8s/internal/aws/sts"
)

const (
	clusterIDHeader  = "x-k8s-aws-id"
	v1Prefix         = "k8s-aws-v1."
	expireHeader     = "X-Amz-Expires"
	expireHeaderTime = "60"
)

// STSTokenGenerator generates a token.
type STSTokenGenerator struct {
	stsClient awssts.PresignClientInterface
}

// NewSTSTokenGenerator creates a new generator.
func NewSTSTokenGenerator(stsClient awssts.PresignClientInterface) *STSTokenGenerator {
	return &STSTokenGenerator{
		stsClient: stsClient,
	}
}

// GenerateToken returns a token valid for clusterID using the given STS client.
func (g *STSTokenGenerator) GenerateToken(ctx context.Context, clusterID string) (string, error) {

	// This code is taken from https://github.com/kubernetes-sigs/aws-iam-authenticator/blob/master/pkg/token/token.go
	// and updated to use AWS SDK v2.

	// generate a sts:GetCallerIdentity request and add our custom cluster ID header
	request, err := g.stsClient.PresignGetCallerIdentity(ctx, &sts.GetCallerIdentityInput{}, func(opts *sts.PresignOptions) {
		opts.ClientOptions = []func(*sts.Options){
			sts.WithAPIOptions(
				smithyhttp.AddHeaderValue(clusterIDHeader, clusterID),
				smithyhttp.AddHeaderValue(expireHeader, expireHeaderTime),
			),
		}
	})
	if err != nil {
		return "", err
	}

	token := v1Prefix + base64.RawURLEncoding.EncodeToString([]byte(request.URL))

	return token, err
}
