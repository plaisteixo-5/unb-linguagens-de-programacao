package functional

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/sqs"
)

func createSession(region string, endpoint string) *session.Session {

    cfg := aws.Config{
        Region: aws.String(region),
        Endpoint: aws.String(endpoint),
    }

   return session.Must(session.NewSession(&cfg))

}