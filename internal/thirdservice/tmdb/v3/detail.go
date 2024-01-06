package tmdb

import (
	"fmt"
	"net/url"
)

// https://developer.themoviedb.org/reference/movie-details
func (c *Client) MovieDetail(id int, option MovieDetailOption) (*MovieDetailResp, error) {
	u := url.URL{
		Scheme: "https",
		Host:   ApiHost,
		Path:   fmt.Sprintf("/3/movie/%d", id),
	}
	q := u.Query()
	q.Add("language", option.Language)
	u.RawQuery = q.Encode()
	resp := MovieDetailResp{}
	if err := c.Get(u.String(), &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
