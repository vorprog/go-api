package util

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/service/sts/stsiface"
)

func GetAwsIdentity() (sts.GetCallerIdentityOutput, error) {
	sess := session.New()
	var svc stsiface.STSAPI = sts.New(sess)
	result, err := svc.GetCallerIdentity(&sts.GetCallerIdentityInput{})

	if err != nil {
		return *result, err
	}

	return *result, nil
}

func GetS3File(s3Bucket string, s3Key string) ([]byte, error) {
	sess := session.New()
	var svc s3iface.S3API = s3.New(sess)
	response, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: &s3Bucket,
		Key:    &s3Key,
	})

	if err != nil {
		return nil, err
	}

	var bodyContent []byte
	_, readErr := response.Body.Read(bodyContent)

	if readErr != nil {
		return nil, readErr
	}

	return bodyContent, nil
}
