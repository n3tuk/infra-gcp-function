package main

import (
	"github.com/n3tuk/infra-gcp-function/infra/stacks/bucket"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func main() {
	app := cdktf.NewApp(
		&cdktf.AppConfig{},
	)

	bucket.NewBucket(app, &bucket.BucketOptions{
		Name:    "n3tuk-gcf-assets",
		Project: "n3tuk-learning-67b95f",
	})

	app.Synth()
}
