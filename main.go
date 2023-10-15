package main

import (
	"github.com/Ingrid-Paulino/Projeto-real-hexagonal-com-go/adapter/input/routes"
	"github.com/Ingrid-Paulino/Projeto-real-hexagonal-com-go/configuration/logger"
	"github.com/gin-gonic/gin"
)

func main()  {
	logger.Info("About to start application")
	router := gin.Default()

	routes.InitRoutes(router)

	if err := router.Run(":8080"); err != nil {
		logger.Error("Error trying to init application on port 8080", err)
		return
	}
}