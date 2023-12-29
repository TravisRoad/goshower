package service

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/TravisRoad/goshower/internal/model"
	"github.com/TravisRoad/goshower/internal/thirdservice/bangumi"
)

var (
	once               = &sync.Once{}
	animeSearchService Searcher
)

const (
	BangumiToken     = "BANGUMI_TOKEN"
	BangumiUserAgent = "BANGUMI_USER_AGENT"
)

type Searcher interface {
	Search(query string, page, size int) (*model.SearchResult, error)
}

func GetAnimeSearchService() Searcher {
	if animeSearchService == nil {
		once.Do(func() {
			token, _ := os.LookupEnv(BangumiToken)
			ua, _ := os.LookupEnv(BangumiUserAgent)
			cli := bangumi.NewClient(token, ua)
			bss := &BangumiSearchService{
				Client: cli,
				Type:   bangumi.SubjectTypeAnime,
			}
			animeSearchService = bss
		})
	}
	return animeSearchService
}

type BangumiSearchService struct {
	Client *bangumi.Client
	Type   int
}

func (s *BangumiSearchService) Search(query string, page, size int) (*model.SearchResult, error) {
	maxResult := min(size, 25)
	start := min((page-1)*maxResult, 0)

	searchResult, err := s.Client.Search(query, bangumi.SearchOptions{
		Start:         start,
		MaxResult:     maxResult,
		Type:          s.Type,
		ResponseGroup: "large",
	})
	if err != nil {
		return nil, nil
	}
	sr := &model.SearchResult{}
	sr.Total = searchResult.Results
	for _, item := range searchResult.List {
		searchItem := model.SearchItem{}
		searchItem.Title = item.Name
		searchItem.TitleCN = item.NameCn
		searchItem.Status = 0
		searchItem.StatusText = ""
		// searchItem.Author = []string{item.}
		date, _ := time.Parse("2006-01-02", item.AirDate)
		searchItem.Date = date
		searchItem.Desc = item.Summary
		searchItem.Rating = uint8(min(item.Rating.Score*10, 100.0))
		searchItem.Source = "bangumi"
		searchItem.Pic = item.Images.Common
		searchItem.ID = fmt.Sprint(item.ID)
		sr.Items = append(sr.Items, searchItem)
	}

	return sr, nil
}
