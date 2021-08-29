package main

import (
   "fmt"
   "github.com/aws/aws-sdk-go/aws"
   "github.com/aws/aws-sdk-go/aws/session"
   "github.com/aws/aws-sdk-go/service/sqs"
   "github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

func newSQS() sqsiface.SQSAPI {

   cfg := aws.Config{
      Region: aws.String("sa-east-1"),
      Endpoint: aws.String("http://localhost:4566"),
   }

   sess := session.Must(session.NewSession(&cfg))

   cliSQS := sqs.New(sess)

   return cliSQS

}

func sendMessage(sqsClient sqsiface.SQSAPI, msg, queueURL string) (*sqs.SendMessageOutput, error) {

   sqsMessage := &sqs.SendMessageInput{
      QueueUrl:    aws.String(queueURL),
      MessageBody: aws.String(msg),
   }

   output, err := sqsClient.SendMessage(sqsMessage)
   if err != nil {
      return nil, fmt.Errorf("could not send message to queue %v: %v", queueURL, err)
   }

   return output, nil
}

func main() {
   sqsClient := newSQS()
   sendMessage(sqsClient, "Mensagem de teste 2", "http://localhost:4566/000000000000/fila_trabalho_lp")
}
