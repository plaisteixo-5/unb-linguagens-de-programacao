package functional

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/sqs"
)


type message struct {
    ID     string  `json:"id"`
    Titulo  string  `json:"title"`
    DataCriacao string  `json:"dataCriacao"`
    Origem  string `json:"origem"`
}

func SendMessage(sess *session.Session, queueURL string) error {

/*     var newMessage = message{ID: "1", Titulo: "Mensagem teste 1", DataCriacao: "1/08/2021", Origem: "SQS producer"} */

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
        MessageBody: aws.String("Teste após refatoração"),
        QueueUrl:   aws.String(queueURL),
    })

    if err != nil {
        return err
    }

    return nil
}