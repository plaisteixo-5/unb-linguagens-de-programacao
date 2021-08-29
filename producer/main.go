package main

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/sqs"
    "github.com/MarceloAmorim25/unb-linguagens-de-programacao/producer/functional"
)

func main() {

    region := "sa-east-1"
    endpoint := "http://localhost:4566"
    queueURL := "http://localhost:4566/000000000000/fila_trabalho_lp"

    sess := createSession(region, endpoint)
    err := SendMsg(sess, queueURL)

    if err != nil {
        fmt.Println("Got an error sending the message:")
        fmt.Println(err)
        return
    }

    fmt.Println("Sent message to queue ")

}

