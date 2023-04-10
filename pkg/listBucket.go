package Bucket

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

var client *s3.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-west-1"))
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	client = s3.NewFromConfig(cfg)
}

type AwsBucket struct {
	Name string `json:"BucketName"`
}

func ListBucketContents(bn AwsBucket) ([]string, error) {
	result, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bn.Name),
	})

	var contents []types.Object
	if err != nil {
		log.Printf("Couldn't list objects in bucket %v. Here's why: %v\n", bn.Name, err)
	} else {
		contents = result.Contents
	}

	var details []string
	for _, c := range contents {
		details = append(details, *c.Key)
	}

	return details, nil
}
