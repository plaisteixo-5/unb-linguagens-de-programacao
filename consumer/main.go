package main

import (
    "fmt"
    "time"
    "github.com/aws/aws-sdk-go/service/sqs"
    "estudos.unb.lp.go.consumer/messages/functional"
)

func pollSqs(chn chan<- *sqs.Message) {

  for {

    fmt.Println("#############################")
    fmt.Println("INICIA LOOP PARA BUSCAR MENSAGENS DA FILA")
    fmt.Println("#############################")

     region := "sa-east-1"
     endpoint := "http://localhost:4566"
     queueURL := "http://localhost:4566/000000000000/fila_trabalho_lp"

     sess := functional.CreateAwsSession(region, endpoint)
     msgResult, err := functional.GetMessages(sess, queueURL, 15)

     if err != nil {
        fmt.Println("Got an error receiving messages:")
        fmt.Println(err)
        return
     }

     fmt.Println("#############################")
     fmt.Println("FINALIZA LOOP PARA BUSCAR MENSAGENS DA FILA")
     fmt.Println("#############################")

    for _, message := range msgResult.Messages {
      fmt.Println("#############################")
      fmt.Println("[INICIO-LOOP] ENVIA MENSAGEM RECEBIDA PARA O CANAL DE MENSAGENS")
      fmt.Println("#############################")
      chn <- message
      fmt.Println("[FIM-LOOP] ENVIA MENSAGEM RECEBIDA PARA O CANAL DE MENSAGENS")
    }

  }

}

func main() {

    fmt.Println("#############################")
    fmt.Println("#############################")
    fmt.Println("ENTRY-POINT: INICIA APLICAÇÃO")
    fmt.Println("#############################")
    fmt.Println("#############################")

     chnMessages := make(chan *sqs.Message, 10)
     go pollSqs(chnMessages)

     for message := range chnMessages {
        fmt.Println("#############################")
        fmt.Println("#############################")
    	fmt.Println("Message ID:     " + *message.MessageId)
        fmt.Println("Message Handle: " + *message.ReceiptHandle)
        fmt.Println("Message Body: " + *message.Body)
        fmt.Println("#############################")
        fmt.Println("#############################")
     }

}
