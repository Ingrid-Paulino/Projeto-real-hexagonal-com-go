package domain

import "time"

type NewsReqDomain struct {
	Subject string
  From time.Time
}

type NewsDomain struct {
	Status string
	TotalResults string
	Articles []Article
}


type Article struct {
	Source string
	Id string
	Name string
	Author string
	Title string
	Description string
	URL string
	URLToImage string
	PublishedAt string
	Content string
}