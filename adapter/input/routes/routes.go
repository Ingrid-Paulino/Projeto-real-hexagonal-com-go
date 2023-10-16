package routes

import (
	"github.com/Ingrid-Paulino/Projeto-real-hexagonal-com-go/adapter/input/controller"
	"github.com/Ingrid-Paulino/Projeto-real-hexagonal-com-go/application/service"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine)  {
	newsService := service.NewNewsService()
	newsController := controller.NewNewsController(newsService)

	r.GET("/news", newsController.GetNews)
}