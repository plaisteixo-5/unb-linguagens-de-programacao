package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type mensagemB struct {
	ID     string  `json:"id"`
	Evento  string  `json:"evento"`
	Destino string  `json:"destino"`
	Duracao  float64 `json:"duracao"`
}

func main() {

	router := gin.Default()

	router.POST("/mensagem-b", salvarMensagemB)

	err := router.Run("localhost:8081")

	if err != nil {
		return
	}

}

func salvarMensagemB(c *gin.Context) {

	var msgB mensagemB

	if err := c.BindJSON(&msgB); err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, msgB)

}
