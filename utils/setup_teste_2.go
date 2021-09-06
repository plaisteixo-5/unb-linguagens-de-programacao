/* package main

import (
    "net/http"
    "github.com/gin-gonic/gin"

    "flag"
    "fmt"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/sqs"
)

type mensagem struct {
    ID     string  `json:"id"`
    Titulo  string  `json:"title"`
    DataCriacao string  `json:"dataCriacao"`
    Origem  string `json:"origem"`
}

var mensagens = []mensagem{
    {ID: "1", Titulo: "Mensagem teste 1", DataCriacao: "1/08/2021", Origem: "SQS producer"},
    {ID: "2", Titulo: "Mensagem teste 2", DataCriacao: "16/08/2021", Origem: "SQS producer"},
    {ID: "3", Titulo: "Mensagem teste 3", DataCriacao: "6/08/2021", Origem: "SQS producer"},
    {ID: "4", Titulo: "Mensagem teste 4", DataCriacao: "2/08/2021", Origem: "SQS producer"},
    {ID: "5", Titulo: "Mensagem teste 5", DataCriacao: "20/08/2021", Origem: "SQS producer"},
    {ID: "6", Titulo: "Mensagem teste 6", DataCriacao: "26/08/2021", Origem: "SQS producer"},
}

func main() {

    router := gin.Default()
    router.POST("/v1/mensagens", geraMensagem)

    router.Run("localhost:8080")

}

func geraMensagem(c *gin.Context) {

    c.IndentedJSON(http.StatusCreated, mensagens)

} */
