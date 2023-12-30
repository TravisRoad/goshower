package bangumi

import "net/http"

type Client struct {
	Cli       *http.Client
	Token     string
	UserAgent string
}
