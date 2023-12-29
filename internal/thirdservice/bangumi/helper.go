package bangumi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) JsonRequest(url string, method string, dst any) error {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("User-Agent", c.UserAgent)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, dst); err != nil {
		return err
	}
	return nil
}

func (c *Client) Get(url string, dst any) error {
	return c.JsonRequest(url, http.MethodGet, dst)
}

func (c *Client) Post(url string, dst any) error {
	return c.JsonRequest(url, http.MethodPost, dst)
}
