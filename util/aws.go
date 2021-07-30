package util

import (
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
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
