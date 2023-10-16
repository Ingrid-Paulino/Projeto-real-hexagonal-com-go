package service

import (
	"fmt"
	"time"

	"github.com/Ingrid-Paulino/Projeto-real-hexagonal-com-go/application/domain"
	"github.com/Ingrid-Paulino/Projeto-real-hexagonal-com-go/configuration/logger"
	"github.com/Ingrid-Paulino/Projeto-real-hexagonal-com-go/configuration/rest_err"
)

type newsService struct{}

func NewNewsService() *newsService {
	return &newsService{}
}

func (*newsService) GetNewService(subject string, from time.Time) (*domain.NewsDomain, *rest_err.RestErr) {
	logger.Info(
		fmt.Sprintf("Init getNewsService function, subject=%s, from=%s", subject, from))

		return nil, nil
}