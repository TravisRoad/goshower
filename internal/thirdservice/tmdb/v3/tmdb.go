package tmdb

import (
	"net/http"
)

type Client struct {
	Cli   *http.Client
	Token string
}
