package main

import (
	"fmt"
	"github.com/MarceloAmorim25/unb-linguagens-de-programacao/producer/functional"
)

const (
	region   = "sa-east-1"
	endpoint = "http://localhost:4566"
	queueURL = "http://localhost:4566/000000000000/fila_trabalho_lp"
)

func main() {

	sess := functional.CreateAwsSession(region, endpoint)

	err := functional.SendMessage(sess, queueURL)

	if err != nil {
		fmt.Println("Got an error sending the message:")
		fmt.Println(err)
		return
	}

}
