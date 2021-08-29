package functional

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
)

func CreateAwsSession(region string, endpoint string) *session.Session {

    cfg := aws.Config{
        Region: aws.String(region),
        Endpoint: aws.String(endpoint),
    }

   return session.Must(session.NewSession(&cfg))

}