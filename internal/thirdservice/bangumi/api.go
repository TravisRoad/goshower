package bangumi

import (
	"context"
	"fmt"
	"net/url"
)

func (c *Client) SearchWithContext(ctx context.Context, query string, options SearchOption) (*SearchResp, error) {
	u := url.URL{
		Scheme: "https",
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

	searchResp := SearchResp{}
	if err := c.GetWithContext(ctx, url, &searchResp); err != nil {
		return nil, err
	}

	return &searchResp, nil
}
func (c *Client) Search(query string, options SearchOption) (*SearchResp, error) {
	u := url.URL{
		Scheme: "https",
		Host:   ApiHost,
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

	searchResp := SearchResp{}
	if err := c.Get(url, &searchResp); err != nil {
		return nil, err
	}

	return &searchResp, nil
}

func (c *Client) GetSubjectDetailWithContext(ctx context.Context, id int) (*SubjectDetail, error) {
	u := url.URL{
		Scheme: "https",
		Host:   ApiHost,
		Path:   fmt.Sprintf("/v0/subjects/%d", id),
	}
	url := u.String()
	subjectDetail := SubjectDetail{}
	if err := c.GetWithContext(ctx, url, &subjectDetail); err != nil {
		return nil, err
	}
	return &subjectDetail, nil
}

func (c *Client) GetSubjectDetail(id int) (*SubjectDetail, error) {
	u := url.URL{
		Scheme: "https",
		Host:   ApiHost,
		Path:   fmt.Sprintf("/v0/subjects/%d", id),
	}
	url := u.String()
	subjectDetail := SubjectDetail{}
	if err := c.Get(url, &subjectDetail); err != nil {
		return nil, err
	}
	return &subjectDetail, nil
}
