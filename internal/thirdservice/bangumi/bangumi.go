package bangumi

type Client struct {
	Token     string
	UserAgent string
}

func NewClient(token string, userAgent string) *Client {
	return &Client{
		Token:     token,
		UserAgent: userAgent,
	}
}
