package functional

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/sqs"
)


func GetMessages(sess *session.Session, queueURL string, timeout int64) (*sqs.ReceiveMessageOutput, error) {

    svc := sqs.New(sess)

    msgResult, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
        AttributeNames: []*string{
            aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
        },
        MessageAttributeNames: []*string{
            aws.String(sqs.QueueAttributeNameAll),
        },
        QueueUrl:            aws.String(queueURL),
        MaxNumberOfMessages: aws.Int64(10),
        VisibilityTimeout:   aws.Int64(timeout),
    })

    if err != nil {
        return nil, err
    }

    return msgResult, nil
}