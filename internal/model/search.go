package model

import "time"

type SearchResult struct {
	Total int          `json:"total"`
	Items []SearchItem `json:"items"`
}

type SearchItem struct {
	ID         string    `json:"id"`
	Source     string    `json:"source"`
	Title      string    `json:"title"`
	TitleCN    string    `json:"title_cn"`
	Desc       string    `json:"description"`
	Date       time.Time `json:"date"`
	Author     []string  `json:"author"`
	Rating     uint8     `json:"rating"`
	Status     uint8     `json:"status"`
	StatusText string    `json:"status_text"`
	Pic        string    `json:"pic"`
}
