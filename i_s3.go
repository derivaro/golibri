package golibri

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func S3GetList(bucket string, svc *s3.S3, params *s3.ListObjectsV2Input) []string {
	fmt.Printf(" 🔍 Loading... ")
	var listUrl []string
	truncatedListing := true
	for truncatedListing {
		resp, err := svc.ListObjectsV2(params)
		if err != nil {
			exitErrf("Unable to list items in %q, %v", bucket, err)
		}
		for _, key := range resp.Contents {

			listUrl = append(listUrl, *key.Key)
		}
		params.ContinuationToken = resp.NextContinuationToken
		truncatedListing = *resp.IsTruncated
	}
	return listUrl
}

func exitErrf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

// --------------------------------------------------------------------------------------------

func S3Del(client *s3.S3, bucket string, key string) (err error) {
	request := &s3.DeleteObjectInput{
		Bucket: &bucket,
		Key:    aws.String(key),
	}
	_, err = client.DeleteObject(request)
	if err != nil {
		return err
	}
	return nil
}
