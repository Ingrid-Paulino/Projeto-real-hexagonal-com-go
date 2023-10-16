package input

import (
	"time"

	"github.com/Ingrid-Paulino/Projeto-real-hexagonal-com-go/application/domain"
	"github.com/Ingrid-Paulino/Projeto-real-hexagonal-com-go/configuration/rest_err"
)

type NewsUseCase interface {
	GetNewService(subject string, from time.Time) (*domain.NewsDomain, *rest_err.RestErr)
}