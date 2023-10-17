package routes

import (
	"github.com/Ingrid-Paulino/Projeto-real-hexagonal-com-go/adapter/input/controller"
	"github.com/Ingrid-Paulino/Projeto-real-hexagonal-com-go/adapter/output/news_http"
	"github.com/Ingrid-Paulino/Projeto-real-hexagonal-com-go/application/service"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine)  {
	newsClient := news_http.NewNewsClient()
	newsService := service.NewNewsService(newsClient)
	newsController := controller.NewNewsController(newsService)

	r.GET("/news", newsController.GetNews)
}