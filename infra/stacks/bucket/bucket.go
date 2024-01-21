package bucket

import (
	"fmt"

	g "github.com/n3tuk/infra-gcp-function/infra/generated/google/provider"
	sb "github.com/n3tuk/infra-gcp-function/infra/generated/google/storagebucket"
	ri "github.com/n3tuk/infra-gcp-function/infra/generated/random/id"
	r "github.com/n3tuk/infra-gcp-function/infra/generated/random/provider"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BucketOptions struct {
	Name                   string
	Project                string
	Location               string `default:"US"`
	ForceDestroy           bool   `default:"false"`
	PublicAccessPrevention string `default:"enforced"`
}

//nolint:gomnd // This is expected to be a fixed value
var (
	randomIDLength        = float64(8)
	bucketNameDescription = "The name of the Google Storage Bucket"
)

func NewBucket(c constructs.Construct, o *BucketOptions) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(c, toResourceName(o, "stack"))

	r.NewRandomProvider(stack,
		toName("random"),
		&r.RandomProviderConfig{},
	)

	g.NewGoogleProvider(stack,
		toName("google"),
		&g.GoogleProviderConfig{},
	)

	// Create and then append a random ID to the name of the bucket itself (but
	// not the name of the bucket resource within the Stack) for any bucket
	// creation to ensure it is globally unique, even across runs of this Stack in
	// the same project and the same region.
	random := ri.NewId(stack,
		toResourceName(o, "random"),
		&ri.IdConfig{
			ByteLength: &randomIDLength,
		},
	)

	bucketName := fmt.Sprintf("%s-%s", o.Name, *random.Hex())
	bucket := sb.NewStorageBucket(stack,
		toResourceName(o, "bucket"),
		&sb.StorageBucketConfig{
			Name:                     &bucketName,
			Project:                  &o.Project,
			Location:                 &o.Location,
			UniformBucketLevelAccess: true,
			PublicAccessPrevention:   &o.PublicAccessPrevention,
		},
	)

	cdktf.NewTerraformOutput(stack,
		toName("bucket_name"),
		&cdktf.TerraformOutputConfig{
			Value:       bucket.Name(),
			Description: &bucketNameDescription,
		},
	)

	return stack
}
