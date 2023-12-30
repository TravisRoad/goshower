package bangumi

import "encoding/json"

/*
https://bangumi.github.io/api/#model-Legacy_SubjectType
subject_type
1 = book
2 = anime
3 = music
4 = game
6 = real
*/
type SubjectType int

const (
	SubjectTypeBook  SubjectType = 1
	SubjectTypeAnime SubjectType = 2
	SubjectTypeMusic SubjectType = 3
	SubjectTypeGame  SubjectType = 4
	SubjectTypeReal  SubjectType = 6
)

const (
	ResGroupSmall  = "small"
	ResGroupMedium = "medium"
	ResGroupLarge  = "large"
)

type SubjectDetail struct {
	ID       int    `json:"id"`
	Type     int    `json:"type"`
	Name     string `json:"name"`
	NameCn   string `json:"name_cn"`
	Summary  string `json:"summary"`
	Nsfw     bool   `json:"nsfw"`
	Locked   bool   `json:"locked"`
	Date     string `json:"date"`
	Platform string `json:"platform"`
	Images   struct {
		Large  string `json:"large"`
		Common string `json:"common"`
		Medium string `json:"medium"`
		Small  string `json:"small"`
		Grid   string `json:"grid"`
	} `json:"images"`
	Infobox []struct {
		Key   string          `json:"key"`
		Value json.RawMessage `json:"value"`
	} `json:"infobox"`
	Volumes       int `json:"volumes"`
	Eps           int `json:"eps"`
	TotalEpisodes int `json:"total_episodes"`
	Rating        struct {
		Rank  int         `json:"rank"`
		Total int         `json:"total"`
		Count map[int]int `json:"count"`
		Score int         `json:"score"`
	} `json:"rating"`
	Collection struct {
		Wish    int `json:"wish"`
		Collect int `json:"collect"`
		Doing   int `json:"doing"`
		OnHold  int `json:"on_hold"`
		Dropped int `json:"dropped"`
	} `json:"collection"`
	Tags []struct {
		Name  string `json:"name"`
		Count int    `json:"count"`
	} `json:"tags"`
}

type SearchItem struct {
	ID         int    `json:"id"`
	URL        string `json:"url"`
	Type       int    `json:"type"`
	Name       string `json:"name"`
	NameCn     string `json:"name_cn"`
	Summary    string `json:"summary"`
	AirDate    string `json:"air_date"`
	AirWeekday int    `json:"air_weekday"`
	Rating     struct {
		Total int         `json:"total"`
		Count map[int]int `json:"count"`
		Score float64     `json:"score"`
	} `json:"rating"`
	Rank   int `json:"rank"`
	Images struct {
		Large  string `json:"large"`
		Common string `json:"common"`
		Medium string `json:"medium"`
		Small  string `json:"small"`
		Grid   string `json:"grid"`
	} `json:"images"`
	Collection struct {
		Wish    int `json:"wish"`
		Collect int `json:"collect"`
		Doing   int `json:"doing"`
		OnHold  int `json:"on_hold"`
		Dropped int `json:"dropped"`
	} `json:"collection"`
}

type SearchResp struct {
	Results int          `json:"results"`
	List    []SearchItem `json:"list"`
}

type SearchOption struct {
	Start         int
	MaxResult     int
	Type          int
	ResponseGroup string
}
