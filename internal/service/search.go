package service

import (
	"net/http"
	"os"
	"sync"

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
			bangumiSearcher.Searcher = &BangumiService{
				Client: cli,
				Type:   toBangumiType(t),
				Source: global.SourceBangumi,
			}
		})
	}
	return bangumiSearcher.Searcher
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
			s := &TMDBService{
				Client: cli,
				Source: global.SourceTMDB,
			}
			tmdbSearcher.Searcher = s
		})
	}
	return tmdbSearcher.Searcher
}
