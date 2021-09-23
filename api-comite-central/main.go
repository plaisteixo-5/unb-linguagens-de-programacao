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

type mensagemB struct {
	ID     string  `json:"id"`
	Evento  string  `json:"evento"`
	Destino string  `json:"destino"`
	Duracao  float64 `json:"duracao"`
}

type mensagemC struct {
	ID     string  `json:"id"`
	Evento  string  `json:"evento"`
	Destino string  `json:"destino"`
	Duracao  float64 `json:"duracao"`
}

func main() {

	router := gin.Default()

	router.POST("/mensagem-a", salvarMensagemA)
	router.POST("/mensagem-b", salvarMensagemB)
	router.POST("/mensagem-c", salvarMensagemC)

	err := router.Run("localhost:8083")

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

func salvarMensagemB(c *gin.Context) {

	var msgB mensagemB

	if err := c.BindJSON(&msgB); err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, msgB)

}

func salvarMensagemC(c *gin.Context) {

	var msgC mensagemC

	if err := c.BindJSON(&msgC); err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, msgC)

}
