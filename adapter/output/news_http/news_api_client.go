package news_http

import (
	"github.com/Ingrid-Paulino/Projeto-real-hexagonal-com-go/adapter/output/model/response"
	"github.com/Ingrid-Paulino/Projeto-real-hexagonal-com-go/application/domain"
	"github.com/Ingrid-Paulino/Projeto-real-hexagonal-com-go/configuration/env"
	"github.com/Ingrid-Paulino/Projeto-real-hexagonal-com-go/configuration/rest_err"
	"github.com/go-resty/resty/v2"
	"github.com/jinzhu/copier"
)

type newsClient struct {}

func NewNewsClient() *newsClient  {
	client = resty.New().SetBaseURL("https//newsapi.org/v2")
	return &newsClient{} 
}

var (
	client *resty.Client
)

func (nc *newsClient) GetNewsPort(newsDomain domain.NewsReqDomain) (*domain.NewsDomain, *rest_err.RestErr) {

	//Faz Requisicao
// q=tesla&from=2023-09-15&apiKey=5d71edc32a3a4113a7248e9da361ed48
	//bibioteca rest ajuda para fazermos requisicoes
	//https://github.com/go-resty/resty

	newsResponse := &response.NewsClientResponse{}

_, err := client.R().
      SetQueryParams(map[string]string{
          "q": newsDomain.Subject,
          "from": newsDomain.From.String(),
          "apiKey": env.GetNewsTokenAPI(),
      }).
			SetResult(newsResponse). //faz o parse do objeto
      Get("/everything")


if err != nil {
	return nil, rest_err.NewInternalServerError("Error trying to call NewsAPI with params")
}

newsResponseDomain := &domain.NewsDomain{}
copier.Copy(newsResponseDomain, newsResponse) //lib que faz copia de objetos. faz a copia da resposta para o nosso domain

	return newsResponseDomain, nil
}