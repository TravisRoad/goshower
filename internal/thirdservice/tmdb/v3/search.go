package tmdb

import (
	"fmt"
	"net/url"
)

// https://developer.themoviedb.org/reference/search-movie
func (c *Client) SearchMovie(query string, option SearchMovieOption) (*SearchMovieResp, error) {
	u := url.URL{
		Scheme: "https",
		Host:   ApiHost,
		Path:   "/3/search/movie",
	}
	q := u.Query()
	q.Add("include_adult", fmt.Sprint(option.Adult))
	q.Add("language", option.Language)
	q.Add("page", fmt.Sprint(option.Page))
	q.Add("query", query)
	q.Add("region", option.Region)
	if option.Year != 0 {
		q.Add("year", fmt.Sprint(option.Year))
	}
	u.RawQuery = q.Encode()

	res := SearchMovieResp{}
	if err := c.Get(u.String(), &res); err != nil {
		return nil, err
	}
	return &res, nil
}
