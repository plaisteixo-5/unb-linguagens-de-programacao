
package main

import (
    "flag"
    "fmt"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/sqs"
)


func GetMessages(sess *session.Session, queueURL *string, timeout *int64) (*sqs.ReceiveMessageOutput, error) {
    // Create an SQS service client
    svc := sqs.New(sess)


    msgResult, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
        AttributeNames: []*string{
            aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
        },
        MessageAttributeNames: []*string{
            aws.String(sqs.QueueAttributeNameAll),
        },
        QueueUrl:            queueURL,
        MaxNumberOfMessages: aws.Int64(1),
        VisibilityTimeout:   timeout,
    })

    if err != nil {
        return nil, err
    }

    return msgResult, nil
}

func main() {

    timeout := flag.Int64("t", 15, "How long, in seconds, that the message is hidden from others")
    flag.Parse()

    sess := session.Must(session.NewSessionWithOptions(session.Options{
         SharedConfigState: session.SharedConfigEnable,
    }))

    queueURL := aws.String("http://localhost:4566/000000000000/fila_trabalho_lp")

    msgResult, err := GetMessages(sess, queueURL, timeout)

    if err != nil {
        fmt.Println("Got an error receiving messages:")
        fmt.Println(err)
        return
    }

    fmt.Println("Message ID:     " + *msgResult.Messages[0].MessageId)
    fmt.Println("Message Handle: " + *msgResult.Messages[0].ReceiptHandle)

}
