package response

type NewsClientResponse struct {
	Status string
	TotalResults string
	Articles []ArticleResponse
}

type ArticleResponse struct {
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