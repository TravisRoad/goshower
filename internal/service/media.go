package service

import (
	"net/http"
	"os"
	"sync"

	"github.com/TravisRoad/goshower/global"
	"github.com/TravisRoad/goshower/internal/model"
	"github.com/TravisRoad/goshower/internal/thirdservice/bangumi"
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
)

type MediaService struct{}

func (s *MediaService) AddMedia(id int, source global.Source, t global.Type) error {
	mediaer := getMediaer(source, t)
	err := mediaer.AddMedia(id)
	return err
}

func (s *MediaService) MediaDetail(id int, source global.Source, t global.Type) (*model.Media, error) {
	Mediaer := getMediaer(source, t)
	if Mediaer == nil {
		global.Logger.Error("no such source", zap.String("source", source.String()), zap.String("type", t.String()))
		return nil, ErrNoSuchSource
	}
	detail, err := Mediaer.MediaDetail(id)
	return detail, err
}

func getMediaer(source global.Source, t global.Type) Mediaer {
	switch source {
	case global.SourceBangumi:
		return getBangumiMediaer(t)
	default:
		return nil
	}
}

// ------------------------
// Mediaer for MediaService
// detail addMedia
// ------------------------

// Bangumi Mediaer
func getBangumiMediaer(t global.Type) Mediaer {
	if bangumiMediaer.Mediaer == nil {
		bangumiMediaer.Once.Do(func() {
			token, _ := os.LookupEnv(BangumiToken)
			ua, _ := os.LookupEnv(BangumiUserAgent)
			cli := &bangumi.Client{
				Cli:       &http.Client{},
				Token:     token,
				UserAgent: ua,
			}
			bangumiMediaer.Mediaer = &BangumiMediaer{
				Client: cli,
			}
		})
	}
	return bangumiMediaer.Mediaer
}

type BangumiMediaer struct {
	Client *bangumi.Client
}

func (s *BangumiMediaer) MediaDetail(id int) (*model.Media, error) {
	return nil, nil
}

func (s *BangumiMediaer) AddMedia(id int) error {
	return nil
}
