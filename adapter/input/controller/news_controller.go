package controller

import "github.com/gin-gonic/gin"

type newsController struct {

}

func NewNewsController() *newsController {
	return &newsController{}
}

func (nc *newsController) GetNews(c *gin.Context)  { // Vai receber a requisição
	//q=tesla&from=2023-09-15&apiKey=5d71edc32a3a4113a7248e9da361ed48
}