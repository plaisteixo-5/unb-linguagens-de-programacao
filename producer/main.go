package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
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
    router.GET("/mensagens", buscaMensagens)
    router.GET("/mensagens/:id", buscaMensagemPorId)

    router.Run("localhost:8080")

}

func buscaMensagens(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, mensagens)
}

func buscaMensagemPorId(c *gin.Context) {

    id := c.Param("id")

    for _, a := range mensagens {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }

    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "messagem não encontrada"})

}