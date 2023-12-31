package input

import (

	"github.com/Ingrid-Paulino/Projeto-real-hexagonal-com-go/application/domain"
	"github.com/Ingrid-Paulino/Projeto-real-hexagonal-com-go/configuration/rest_err"
)

type NewsUseCase interface {
	GetNewService(domain.NewsReqDomain) (*domain.NewsDomain, *rest_err.RestErr)
}