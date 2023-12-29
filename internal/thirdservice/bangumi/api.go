package bangumi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) Search(query string, options SearchOptions) (SearchResp, error) {
	searchResp := SearchResp{}
	u := url.URL{
		Scheme: "https",
		Host:   "api.bgm.tv",
		Path:   fmt.Sprintf("/search/subject/%s", query),
	}
	q := u.Query()
	if options.Type != 0 {
		q.Add("type", fmt.Sprintf("%d", options.Type))
	}
	if options.Start != 0 {
		q.Add("start", fmt.Sprintf("%d", options.Start))
	}
	if options.MaxResult != 0 {
		q.Add("max_results", fmt.Sprintf("%d", options.MaxResult))
	}
	if len(options.ResponseGroup) != 0 {
		q.Add("responseGroup", options.ResponseGroup)
	}
	u.RawQuery = q.Encode()
	url := u.String()

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return searchResp, err
	}

	if len(c.UserAgent) != 0 {
		req.Header.Add("User-Agent", c.UserAgent)
	}
	if len(c.Token) != 0 {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	}

	resp, err := client.Do(req)
	if err != nil {
		return searchResp, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return searchResp, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return searchResp, err
	}
	if err := json.Unmarshal(body, &searchResp); err != nil {
		return searchResp, err
	}

	return searchResp, nil
}
