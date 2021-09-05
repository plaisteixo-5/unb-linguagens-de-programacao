package main

import (
    "fmt"
    "estudos.unb.lp.go.producer/messages/functional"
)

func main() {

    region := "sa-east-1"
    endpoint := "http://localhost:4566"
    queueURL := "http://localhost:4566/000000000000/fila_trabalho_lp"

    sess := functional.CreateAwsSession(region, endpoint)

    err := functional.SendMessage(sess, queueURL)

    if err != nil {
        fmt.Println("Got an error sending the message:")
        fmt.Println(err)
        return
    }

}

