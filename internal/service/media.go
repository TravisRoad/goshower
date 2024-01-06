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

type MediaSingleton struct {
	Once    *sync.Once
	Mediaer Mediaer
}

func NewMediaSingleton() MediaSingleton {
	return MediaSingleton{
		Once:    &sync.Once{},
		Mediaer: nil,
	}
}

var (
	bangumiMediaer = NewMediaSingleton()
	tmdbMediaer    = NewMediaSingleton()
)

type MediaService struct{}

func (s *MediaService) MediaDetail(id int, source global.Source) (*model.Media, error) {
	Mediaer := getMediaer(source)
	if Mediaer == nil {
		global.Logger.Error("no such source", zap.String("source", source.String()))
		return nil, ErrNoSuchSource
	}
	detail, err := Mediaer.MediaDetail(id)
	return detail, err
}

func getMediaer(source global.Source) Mediaer {
	switch source {
	case global.SourceBangumi:
		return getBangumiMediaer()
	case global.SourceTMDB:
		return getTMDBMediaer()
	default:
		return nil
	}
}

// ------------------------
// Mediaer for MediaService
// detail addMedia
// ------------------------

// Bangumi Mediaer
func getBangumiMediaer() Mediaer {
	if bangumiMediaer.Mediaer == nil {
		bangumiMediaer.Once.Do(func() {
			token, _ := os.LookupEnv(BangumiToken)
			ua, _ := os.LookupEnv(BangumiUserAgent)
			cli := &bangumi.Client{
				Cli:       &http.Client{},
				Token:     token,
				UserAgent: ua,
			}
			bangumiMediaer.Mediaer = &BangumiService{
				Source: global.SourceBangumi,
				Client: cli,
			}
		})
	}
	return bangumiMediaer.Mediaer
}

// tmdb mediaer
func getTMDBMediaer() Mediaer {
	if tmdbMediaer.Mediaer == nil {
		tmdbMediaer.Once.Do(func() {
			token, _ := os.LookupEnv(TMDBToken)
			cli := &tmdb.Client{
				Cli:   &http.Client{},
				Token: token,
			}
			tmdbMediaer.Mediaer = &TMDBService{
				Source: global.SourceTMDB,
				Client: cli,
			}
		})
	}
	return tmdbMediaer.Mediaer
}
