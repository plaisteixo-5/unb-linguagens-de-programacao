package main

import (
	"fmt"
	"time"

	"estudos.unb.lp.go.consumer/messages/functional"
	"github.com/aws/aws-sdk-go/service/sqs"
)

const (
	region   = "sa-east-1"
	endpoint = "http://localhost:4566"
)

func main() {

	fmt.Println("")
	fmt.Println("#############################")
	fmt.Println("[MAIN] ENTRY-POINT: INICIA APLICAÇÃO")
	fmt.Println("#############################")
	fmt.Println("")

	fmt.Println("[MAIN] CHANNEL -> RECEBE MENSAGENS...")
	messagesChannel := make(chan *sqs.Message, 10)

	fmt.Println("[MAIN] Loading...")
	time.Sleep(5 * time.Second)

	sess := functional.CreateAwsSession(region, endpoint)

	fmt.Println("[MAIN] POLLING DA FILA...")

	// Função executada de forma concorrente...
	go functional.FetchQueueMessages(messagesChannel, sess)

	fmt.Println("[MAIN] QUANTIDADE DE MENSAGENS:", len(messagesChannel))

	functional.HandleAllReceivedMessages(messagesChannel, sess)

}
