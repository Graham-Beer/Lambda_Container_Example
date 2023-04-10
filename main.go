package main

import (
	Bucket "Bucket/pkg"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(Bucket.ListBucketContents)
}
