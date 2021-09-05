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

func pollSqs(chn chan<- *sqs.Message, sess *session.Session) {

  for {

     fmt.Println("[INICIA LOOP - FETCH] PARA BUSCAR MENSAGENS DA FILA")
     fmt.Println("Preparing application...")
     time.Sleep(3 * time.Second)

     fmt.Println("Go...")

     msgResult, err := functional.GetMessages(sess, queueURL, 15)
     fmt.Println("[POLL] QUANTIDADE DE MENSAGENS:", len(msgResult.Messages))

     if err != nil {
        fmt.Println("Got an error receiving messages:")
        fmt.Println(err)
        return
     }

     fmt.Println("#############################")
     fmt.Println("FINALIZA LOOP PARA BUSCAR MENSAGENS DA FILA")
     fmt.Println("#############################")

    for _, message := range msgResult.Messages {
      fmt.Println("[INICIO-LOOP-MENSAGENS RECEBIDAS] ENVIA MENSAGEM RECEBIDA PARA O CANAL DE MENSAGENS")
      chn <- message
      fmt.Println("[FIM-LOOP-MENSAGENS RECEBIDAS] ENVIA MENSAGEM RECEBIDA PARA O CANAL DE MENSAGENS")
    }

  }

}

func handleMessage(message sqs.Message, sess *session.Session, queueURL string, messageCounter int) {

    fmt.Println("[HANDLING MESSAGE] MESSAGE NUMBER=%d BEING SEND TO RESPECTIVE DESTINATION", messageCounter)
    functional.SendByPOSTRequest(message)
    functional.DeleteMessage(sess, queueURL, *message.ReceiptHandle)

}

func main() {

    fmt.Println("#############################")
    fmt.Println("#############################")
    fmt.Println("ENTRY-POINT: INICIA APLICAÇÃO")
    fmt.Println("#############################")
    fmt.Println("#############################")

     chnMessages := make(chan *sqs.Message, 10)
     sess := functional.CreateAwsSession(region, endpoint)

     go pollSqs(chnMessages, sess)

     fmt.Println("[POLL] QUANTIDADE DE MENSAGENS:", len(chnMessages))

     messageCounter := 0
     for message := range chnMessages {

          messageCounter += 1
          go handleMessage(*message, sess, queueURL, messageCounter)

     }

}
