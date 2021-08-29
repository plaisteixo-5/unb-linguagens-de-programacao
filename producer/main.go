package main

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/sqs"
)

type mensagem struct {
    ID     string  `json:"id"`
    Titulo  string  `json:"title"`
    DataCriacao string  `json:"dataCriacao"`
    Origem  string `json:"origem"`
}

func SendMsg(sess *session.Session, queueURL *string) error {

/*     var newMessage = mensagem{ID: "1", Titulo: "Mensagem teste 1", DataCriacao: "1/08/2021", Origem: "SQS producer"} */

    svc := sqs.New(sess)

    _, err := svc.SendMessage(&sqs.SendMessageInput{
        DelaySeconds: aws.Int64(2),
        MessageAttributes: map[string]*sqs.MessageAttributeValue{
            "ID": &sqs.MessageAttributeValue{
                DataType:    aws.String("Number"),
                StringValue: aws.String("1"),
            },
            "Titulo": &sqs.MessageAttributeValue{
                DataType:    aws.String("String"),
                StringValue: aws.String("Mensagem teste 1"),
            },
            "DataCriacao": &sqs.MessageAttributeValue{
                DataType:    aws.String("String"),
                StringValue: aws.String("1/08/2021"),
            },
            "Origem": &sqs.MessageAttributeValue{
                DataType:    aws.String("String"),
                StringValue: aws.String("SQS producer"),
            },
        },
        MessageBody: aws.String("Teste com json"),
        QueueUrl:    queueURL,
    })

    if err != nil {
        return err
    }

    return nil
}

func main() {

    cfg := aws.Config{
          Region: aws.String("sa-east-1"),
          Endpoint: aws.String("http://localhost:4566"),
    }

    sess := session.Must(session.NewSession(&cfg))

    queueURL := aws.String("http://localhost:4566/000000000000/fila_trabalho_lp")

    err := SendMsg(sess, queueURL)
    if err != nil {
        fmt.Println("Got an error sending the message:")
        fmt.Println(err)
        return
    }

    fmt.Println("Sent message to queue ")

}

