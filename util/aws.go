package util

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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

func GetS3File(bucketName string, filename string, region string) (file io.Reader, err error) {
	s3Session, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return nil, err
	}
	s3Client := s3.New(s3Session)
	s3Object, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
	})
	if err != nil {
		return nil, err
	}
	return s3Object.Body, nil
}
