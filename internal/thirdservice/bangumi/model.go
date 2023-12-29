package bangumi

/*
https://bangumi.github.io/api/#model-Legacy_SubjectType
subject_type
1 = book
2 = anime
3 = music
4 = game
6 = real
*/

const (
	SubjectTypeBook  = 1
	SubjectTypeAnime = 2
	SubjectTypeMusic = 3
	SubjectTypeGame  = 4
	SubjectTypeReal  = 6

	ResGroupSmall  = "small"
	ResGroupMedium = "medium"
	ResGroupLarge  = "large"
)

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

type SearchOptions struct {
	Start         int
	MaxResult     int
	Type          int
	ResponseGroup string
}
