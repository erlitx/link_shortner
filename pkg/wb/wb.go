package wb

import (
	"net/http"
	"time"
)

type Config struct {
	WBTokenStats string `envconfig:"WB_TOKEN_STATS"      required:"true"`
	WBTokenSale  string `envconfig:"WB_TOKEN_SALE"      required:"true"`
}


type Client struct {
	Client       *http.Client
	WBTokenStats string
	WBTokenSale string
}

func New(c Config) *Client {
	return &Client{
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
		WBTokenStats: c.WBTokenStats,
		WBTokenSale: c.WBTokenSale,
	}
}
