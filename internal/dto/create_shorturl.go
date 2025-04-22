package dto


type CreateShortUrlOutput struct {
	ShortURL string `json:"shortUrl"`
}

type CreateShortUrlInput struct {
	RawURL string `json:"rawUrl"`
	Host string 
}
