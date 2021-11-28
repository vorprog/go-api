package util

import (
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/service/sts/stsiface"
)

func GetAwsIdentity() sts.GetCallerIdentityOutput {
	sess := session.New()
	var svc stsiface.STSAPI = sts.New(sess)
	result, err := svc.GetCallerIdentity(&sts.GetCallerIdentityInput{})

	if err != nil {
		Log(err)
		os.Exit(1)
	}

	return *result
}

func GetS3File(s3Bucket string, s3Key string) []byte {
	sess := session.New()
	var svc s3iface.S3API = s3.New(sess)
	response, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: &s3Bucket,
		Key:    &s3Key,
	})

	if err != nil {
		Log(err)
		os.Exit(1)
	}

	var bodyContent []byte
	_, readErr := response.Body.Read(bodyContent)

	if readErr != nil {
		Log(readErr)
		os.Exit(1)
	}

	return bodyContent
}
