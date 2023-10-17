package output

import (
	"github.com/Ingrid-Paulino/Projeto-real-hexagonal-com-go/application/domain"
	"github.com/Ingrid-Paulino/Projeto-real-hexagonal-com-go/configuration/rest_err"
)

type NewsPort interface {
	GetNewsPort(domain.NewsReqDomain) (*domain.NewsDomain, *rest_err.RestErr)
}