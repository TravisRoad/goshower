package bangumi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) makeRequestWithContext(ctx context.Context, url string, method string, dst any) error {
	client := c.Cli
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
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

func (c *Client) makeRequest(url string, method string, dst any) error {
	client := c.Cli
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
	return c.makeRequest(url, http.MethodGet, dst)
}

func (c *Client) Post(url string, dst any) error {
	return c.makeRequest(url, http.MethodPost, dst)
}

func (c *Client) GetWithContext(ctx context.Context, url string, dst any) error {
	return c.makeRequestWithContext(ctx, url, http.MethodGet, dst)
}

func (c *Client) PostWithContext(ctx context.Context, url string, dst any) error {
	return c.makeRequestWithContext(ctx, url, http.MethodPost, dst)
}
