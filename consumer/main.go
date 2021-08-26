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

func main() {

    router := gin.Default()
    router.POST("/mensagens", criarMensagem)

    router.Run("localhost:8080")

}

func criarMensagem(c *gin.Context) {

    var novaMensagem mensagem

    if err := c.BindJSON(&novaMensagem); err != nil {
        return
    }

    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)

}