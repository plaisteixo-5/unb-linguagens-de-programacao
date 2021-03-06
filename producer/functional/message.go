package functional

import (
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/sqs"
)


type message struct {
    ID          string  `json:"id"`
    Title      string  `json:"title"`
    CreatedDate string  `json:"created_date"`
    Origin      string `json:"origin"`
}

func SendMessage(sess *session.Session, queueURL string) error {

    svc := sqs.New(sess)

    fmt.Println("Sending new message to queue...")

    _, err := svc.SendMessage(&sqs.SendMessageInput{
        DelaySeconds: aws.Int64(2),
        MessageAttributes: map[string]*sqs.MessageAttributeValue{
            "ID": &sqs.MessageAttributeValue{
                DataType:    aws.String("Number"),
                StringValue: aws.String("1"),
            },
            "Title": &sqs.MessageAttributeValue{
                DataType:    aws.String("String"),
                StringValue: aws.String("Mensagem teste 1"),
            },
            "CreatedDate": &sqs.MessageAttributeValue{
                DataType:    aws.String("String"),
                StringValue: aws.String("1/08/2021"),
            },
            "Origin": &sqs.MessageAttributeValue{
                DataType:    aws.String("String"),
                StringValue: aws.String("SQS producer"),
            },
        },
        MessageBody: aws.String("Teste após refatoração"),
        QueueUrl:   aws.String(queueURL),
    })

    if err != nil {
        return err
    }

    return nil
}

