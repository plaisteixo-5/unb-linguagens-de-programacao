package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type mensagemA struct {
	ID     string  `json:"id"`
	Evento  string  `json:"evento"`
	Destino string  `json:"destino"`
	Duracao  float64 `json:"duracao"`
}

func main() {

	router := gin.Default()

	router.POST("/mensagem-a", salvarMensagemA)

	err := router.Run("localhost:8080")

	if err != nil {
		return
	}

}

func salvarMensagemA(c *gin.Context) {

	var msgA mensagemA

	if err := c.BindJSON(&msgA); err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, msgA)

}


