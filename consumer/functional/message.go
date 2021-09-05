package functional

import (
    "fmt"
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


func DeleteMessage(sess *session.Session, queueURL string, messageHandle string) error {

    svc := sqs.New(sess)

    _, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
        QueueUrl:      aws.String(queueURL),
        ReceiptHandle: aws.String(messageHandle),
    })

    if err != nil {
        return err
    }

    return nil
}

func SendByPOSTRequest(message sqs.Message) {

    fmt.Println("#############################")
    fmt.Println("#############################")
    fmt.Println("Message ID:     " + *message.MessageId)
    fmt.Println("Message Handle: " + *message.ReceiptHandle)
    fmt.Println("Message Body: " + *message.Body)
    fmt.Println("#############################")
    fmt.Println("#############################")

}
