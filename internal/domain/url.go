package domain


type RawURL string

type ShortURL string

type URL struct {
	RawURL   RawURL   `json:"rawUrl"`
	ShortURL ShortURL `json:"shortUrl"`
}

func NewUrl(rawUrl, shortUrl string) (URL, error) {
	var url URL

	if rawUrl == "" {
		return url, ErrEmptyUrl
	}

	url = URL{
		RawURL:   RawURL(rawUrl),
		ShortURL: ShortURL(shortUrl),
	}

	return url, nil
}
