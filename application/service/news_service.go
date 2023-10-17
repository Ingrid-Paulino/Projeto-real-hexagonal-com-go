package service

import (
	"fmt"

	"github.com/Ingrid-Paulino/Projeto-real-hexagonal-com-go/application/domain"
	"github.com/Ingrid-Paulino/Projeto-real-hexagonal-com-go/application/port/output"
	"github.com/Ingrid-Paulino/Projeto-real-hexagonal-com-go/configuration/logger"
	"github.com/Ingrid-Paulino/Projeto-real-hexagonal-com-go/configuration/rest_err"
)

type newsService struct{
	newsPort output.NewsPort
}

func NewNewsService(newsPort output.NewsPort) *newsService {
	return &newsService{newsPort: newsPort}
}

func (ns *newsService) GetNewService(newsDomain domain.NewsReqDomain) (*domain.NewsDomain, *rest_err.RestErr) {
	logger.Info(
		fmt.Sprintf("Init getNewsService function, subject=%s, from=%s", newsDomain.Subject, newsDomain.From))

		newsDomainResponse, err := ns.newsPort.GetNewsPort(newsDomain)
		return newsDomainResponse, err
}