package model

import (
	"time"

	"github.com/TravisRoad/goshower/global"
)

type SearchResult struct {
	TotalResult int          `json:"total_result"`
	TotalPage   int          `json:"total_page"`
	Page        int          `json:"page"`
	Items       []SearchItem `json:"items"`
}

type SearchItem struct {
	ID         string        `json:"id"`
	Source     global.Source `json:"source"`
	Title      string        `json:"title"`
	TitleCN    string        `json:"title_cn"`
	Desc       string        `json:"description"`
	Date       time.Time     `json:"date"`
	Author     []string      `json:"author"`
	Rating     uint8         `json:"rating"`
	Status     uint8         `json:"status"`
	StatusText string        `json:"status_text"`
	Pic        string        `json:"pic"`
}
