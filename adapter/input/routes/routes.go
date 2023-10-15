package routes

import (
	"github.com/Ingrid-Paulino/Projeto-real-hexagonal-com-go/adapter/input/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine)  {
	
	newsController := controller.NewNewsController()

	r.GET("/news", newsController.GetNews)
}