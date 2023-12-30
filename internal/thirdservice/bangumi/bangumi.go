package bangumi

import "net/http"

type Client struct {
	Cli       *http.Client
	Host      string
	Token     string
	UserAgent string
}
