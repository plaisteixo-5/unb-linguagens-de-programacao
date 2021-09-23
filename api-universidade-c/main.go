package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


type mensagemC struct {
	ID     string  `json:"id"`
	Evento  string  `json:"evento"`
	Destino string  `json:"destino"`
	Duracao  float64 `json:"duracao"`
}

func main() {

	router := gin.Default()

	router.POST("/mensagem-c", salvarMensagemC)

	err := router.Run("localhost:8082")

	if err != nil {
		return
	}

}

func salvarMensagemC(c *gin.Context) {

	var msgC mensagemC

	if err := c.BindJSON(&msgC); err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, msgC)

}
