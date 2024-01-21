package bucket_test

import (
	"os"
	"testing"

	"github.com/n3tuk/infra-gcp-function/infra/stacks/bucket"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/stretchr/testify/assert"
)

func TestNewBucket(t *testing.T) {
	tmp, err := os.MkdirTemp(os.TempDir(), "bucket")
	assert.NoError(t, err, "There should be no error creating a temporary directory for CDKTF Stacks")

	// Ensure we remove test stacks and configurations after testing
	defer os.RemoveAll(tmp)

	// Create the Application and Stack for testing
	app := cdktf.Testing_App(&cdktf.TestingAppConfig{Outdir: &tmp})
	stack := bucket.NewBucket(app, &bucket.BucketOptions{Name: "assets"})

	// Run the synthesis of the Stack and validate it
	code := cdktf.Testing_FullSynth(stack)
	assertion := cdktf.Testing_ToBeValidTerraform(code)
	assert.True(t, *assertion, "NewBucket and CDKTF should create valid Terraform")
}
