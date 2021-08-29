package main

import (
    "flag"
    "fmt"
    "estudos.unb.lp.go.consumer/messages/functional"
)

func main() {

    region := "sa-east-1"
    endpoint := "http://localhost:4566"
    queueURL := "http://localhost:4566/000000000000/fila_trabalho_lp"
    timeout := flag.Int64("t", 15, "How long, in seconds, that the message is hidden from others")
    flag.Parse()

    sess := functional.CreateAwsSession(region, endpoint)
    msgResult, err := functional.GetMessages(sess, queueURL, timeout)

    if err != nil {
        fmt.Println("Got an error receiving messages:")
        fmt.Println(err)
        return
    }

    for _, message := range msgResult.Messages {
    	fmt.Println("Message ID:     " + *message.MessageId)
        fmt.Println("Message Handle: " + *message.ReceiptHandle)
        fmt.Println("Message Body: " + *message.Body)
    }

}
