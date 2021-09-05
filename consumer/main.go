package main

import (
    "fmt"
    "time"
    "github.com/aws/aws-sdk-go/service/sqs"
    "github.com/aws/aws-sdk-go/aws/session"
    "estudos.unb.lp.go.consumer/messages/functional"
)


const region = "sa-east-1"
const endpoint = "http://localhost:4566"
const queueURL = "http://localhost:4566/000000000000/fila_trabalho_lp"

func sendMessagesToChannel(chn chan<- *sqs.Message, msgResult *sqs.ReceiveMessageOutput) {

      for _, message := range msgResult.Messages {

        fmt.Println("[POLL SQS] Enviando mensagens para o channel...")
        chn <- message

      }

}

func pollSqs(chn chan<- *sqs.Message, sess *session.Session) {

  for {

     fmt.Println("[POLL SQS] Preparing application to fetch messages...")
     time.Sleep(3 * time.Second)

     fmt.Println("[POLL SQS] Go...")

     msgResult, err := functional.GetMessages(sess, queueURL, 10)

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

func handleMessage(message sqs.Message, sess *session.Session, queueURL string, messageCounter int) {

    fmt.Println("[METHOD - handleMessage] MESSAGE NUMBER= ", messageCounter, " BEING SEND TO RESPECTIVE DESTINATION")
    functional.SendByPOSTRequest(message)
    functional.DeleteMessage(sess, queueURL, *message.ReceiptHandle)

}

func handleAllReceivedMessages(chnMessages chan *sqs.Message, sess *session.Session) {

    messageCounter := 0
    for message := range chnMessages {

       messageCounter += 1

       fmt.Println("[METHOD - handleAllReceivedMessages] As mensagens estão sendo enviadas para seu destino final...")
       go handleMessage(*message, sess, queueURL, messageCounter)

   }

}


func main() {

     fmt.Println("#############################")
     fmt.Println("#############################")
     fmt.Println("[MAIN] ENTRY-POINT: INICIA APLICAÇÃO")
     fmt.Println("#############################")
     fmt.Println("#############################")

     fmt.Println("[MAIN] CHANNEL -> RECEBE MENSAGENS...")
     messagesChannel := make(chan *sqs.Message, 10)

     fmt.Println("[MAIN] Loading...")
     time.Sleep(5 * time.Second)

     sess := functional.CreateAwsSession(region, endpoint)

     fmt.Println("[MAIN] POLLING DA FILA...")

     // Função executada de forma concorrente...
     go pollSqs(messagesChannel, sess)

     fmt.Println("[MAIN] QUANTIDADE DE MENSAGENS:", len(messagesChannel))

     handleAllReceivedMessages(messagesChannel, sess)

}
