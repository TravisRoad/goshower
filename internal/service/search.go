package service

import (
	"math"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/TravisRoad/goshower/global"
	"github.com/TravisRoad/goshower/internal/model"
	"github.com/TravisRoad/goshower/internal/thirdservice/bangumi"
	"github.com/TravisRoad/goshower/internal/thirdservice/tmdb/v3"
	"go.uber.org/zap"
)

type SearchSingleton struct {
	Once     *sync.Once
	Searcher Searcher
}

func NewSearchSingleton() SearchSingleton {
	return SearchSingleton{
		Once:     &sync.Once{},
		Searcher: nil,
	}
}

var (
	bangumiSearcher = NewSearchSingleton()
	tmdbSearcher    = NewSearchSingleton()
)

const (
	BangumiToken     = "BANGUMI_TOKEN"
	BangumiUserAgent = "BANGUMI_USER_AGENT"

	TMDBToken = "TMDB_TOKEN"
)

type SearchService struct{}

func (s *SearchService) Search(query string, page, size int, source global.Source, t global.Type) (*model.SearchResult, error) {
	searcher := GetSearcher(source, t)
	if searcher == nil {
		global.Logger.Error("no such source", zap.String("source", source.String()), zap.String("type", t.String()))
		return nil, ErrNoSuchSource
	}
	return searcher.Search(query, page, size)
}

// Searcher
// Search the query thing

func GetSearcher(source global.Source, t global.Type) Searcher {
	switch source {
	case global.SourceBangumi:
		return getBangumiSearcher(t)
	case global.SourceTMDB:
		return getTMDBSearcher(t)
	default:
		return nil
	}
}

// Bangumi Searcher

func getBangumiSearcher(t global.Type) Searcher {
	if bangumiSearcher.Searcher == nil {
		bangumiSearcher.Once.Do(func() {
			token, _ := os.LookupEnv(BangumiToken)
			ua, _ := os.LookupEnv(BangumiUserAgent)
			cli := &bangumi.Client{
				Cli:       &http.Client{},
				Token:     token,
				UserAgent: ua,
			}
			bss := &BangumiSearcher{
				Client: cli,
				Type:   int(toBangumiType(t)),
				Source: global.SourceBangumi,
			}
			bangumiSearcher.Searcher = bss
		})
	}
	return bangumiSearcher.Searcher
}

type BangumiSearcher struct {
	Client *bangumi.Client
	Type   int
	Source global.Source
}

func (s *BangumiSearcher) Search(query string, page, size int) (*model.SearchResult, error) {
	maxResult := min(size, 25)
	start := min((page-1)*maxResult, 0)

	searchResult, err := s.Client.Search(query, bangumi.SearchOption{
		Start:         start,
		MaxResult:     maxResult,
		Type:          s.Type,
		ResponseGroup: "large",
	})
	if err != nil {
		return nil, nil
	}
	sr := &model.SearchResult{}
	sr.TotalResult = searchResult.Results
	sr.Page = page
	sr.TotalPage = int(math.Ceil(float64(searchResult.Results) / float64(maxResult)))
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
		searchItem.Source = s.Source
		searchItem.Pic = item.Images.Common

		id, err := global.Sqids.Encode([]uint64{uint64(s.Source), uint64(item.ID)})
		if err != nil {
			return nil, err
		}

		searchItem.ID = id
		sr.Items = append(sr.Items, searchItem)
	}

	return sr, nil
}

// TMDB Searcher

func getTMDBSearcher(t global.Type) Searcher {
	if tmdbSearcher.Searcher == nil {
		tmdbSearcher.Once.Do(func() {
			token, _ := os.LookupEnv(TMDBToken)
			cli := &tmdb.Client{
				Cli:   &http.Client{},
				Token: token,
			}
			s := &TMDBSearcher{
				Client: cli,
				Source: global.SourceTMDB,
			}
			tmdbSearcher.Searcher = s
		})
	}
	return tmdbSearcher.Searcher
}

type TMDBSearcher struct {
	Client *tmdb.Client
	Source global.Source
}

func (s *TMDBSearcher) Search(query string, page, size int) (*model.SearchResult, error) {
	tmdbResult, err := s.Client.SearchMovie(query, tmdb.SearchMovieOption{Language: "zh-CN", Page: page})
	if err != nil {
		return nil, err
	}
	sr := &model.SearchResult{}
	sr.TotalPage = tmdbResult.TotalPages
	sr.TotalResult = tmdbResult.TotalResults
	sr.Page = tmdbResult.Page
	sr.Items = make([]model.SearchItem, len(tmdbResult.Results))
	for i, item := range tmdbResult.Results {
		searchItem := model.SearchItem{}
		searchItem.Title = item.Title
		searchItem.TitleCN = item.OriginalTitle
		searchItem.Status = 0
		searchItem.StatusText = ""
		date, _ := time.Parse("2006-01-02", item.ReleaseDate)
		searchItem.Date = date
		searchItem.Desc = item.Overview
		searchItem.Rating = uint8(min(item.VoteAverage*10, 100.0))
		searchItem.Source = s.Source
		searchItem.Pic = item.PosterPath
		id, err := global.Sqids.Encode([]uint64{uint64(s.Source), uint64(item.ID)})
		if err != nil {
			return nil, err
		}
		searchItem.ID = id
		sr.Items[i] = searchItem
	}

	return nil, nil
}
