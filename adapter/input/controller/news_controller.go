package controller

import (
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/adapter/input/model/request"
	"github.com/Ingrid-Paulino/Projeto-real-hexagonal-com-go/application/domain"
	"github.com/Ingrid-Paulino/Projeto-real-hexagonal-com-go/application/port/input"
	"github.com/Ingrid-Paulino/Projeto-real-hexagonal-com-go/configuration/logger"
	"github.com/Ingrid-Paulino/Projeto-real-hexagonal-com-go/configuration/validation"
	"github.com/gin-gonic/gin"
)

type newsController struct {
	newsUseCase input.NewsUseCase
}

func NewNewsController(newsUseCase input.NewsUseCase) *newsController {
	return &newsController{}
}

func (nc *newsController) GetNews(c *gin.Context)  { // Vai receber a requisição

	logger.Info("Init GetNews controller api")
	newsRequest := request.NewsRequest{}

	if err := c.ShouldBindQuery(&newsRequest); err != nil {
		logger.Error("Error trying to validate field from request", err)
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	newsDomain := domain.NewsReqDomain{Subject: newsRequest.Subject, From: newsRequest.From}

	_, _ = nc.newsUseCase.GetNewService(newsDomain)
}