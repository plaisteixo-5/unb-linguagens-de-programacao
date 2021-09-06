package functional

import (
    "fmt"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/sqs"
    "time"
)

const (
    queueURL = "http://localhost:4566/000000000000/fila_trabalho_lp"
)


func FetchQueueMessages(chn chan<- *sqs.Message, sess *session.Session) {

  for {

     fmt.Println("[POLL SQS] Preparing application to fetch messages...")
     time.Sleep(3 * time.Second)

     fmt.Println("[POLL SQS] Go...")

     msgResult, err := GetMessages(sess, queueURL, 10)

     fmt.Println("[POLL SQS] Operation to get message was done:")
     fmt.Println("[POLL SQS] QUANTIDADE DE MENSAGENS:", len(msgResult.Messages))
     time.Sleep(3 * time.Second)

     if err != nil {
        fmt.Println("[POLL SQS] Got an error receiving messages:")
        fmt.Println(err)
        return
     }

    sendMessagesToChannel(chn, msgResult)

    time.Sleep(3 * time.Second)
    fmt.Println("[POLL SQS] Preparing application for polling again...")

  }

}