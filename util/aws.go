package util

import (
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/service/sts/stsiface"
)

func GetAwsIdentity() sts.GetCallerIdentityOutput {
	sess := session.New()
	var svc stsiface.STSAPI = sts.New(sess)
	result, err := svc.GetCallerIdentity(&sts.GetCallerIdentityInput{})

	if err != nil {
		log.Fatal(err)
	}

	return *result
}
