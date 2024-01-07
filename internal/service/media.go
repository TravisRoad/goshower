package service

import (
	"errors"
	"net/http"
	"os"
	"sync"

	"github.com/TravisRoad/goshower/global"
	"github.com/TravisRoad/goshower/internal/model"
	"github.com/TravisRoad/goshower/internal/thirdservice/bangumi"
	"github.com/TravisRoad/goshower/internal/thirdservice/tmdb/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
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
	var media *model.Media
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		// check if the media item exsits in database
		item := model.Media{}
		err := tx.Model(&model.Media{}).Where("media_id = ? AND source = ?", id, source).Take(&item).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// if not exsit, request media detail
			Mediaer, err := getMediaer(source)
			if err != nil {
				global.Logger.Error("no such source", zap.Any("source", source), zap.Error(err))
				return ErrNoSuchSource
			}
			detail, err := Mediaer.MediaDetail(id)
			if err != nil {
				return err
			}
			media = detail
			return nil
		}
		if err != nil {
			return err
		}
		media = &item
		return nil
	})
	return media, err
}

func getMediaer(source global.Source) (Mediaer, error) {
	switch source {
	case global.SourceBangumi:
		return getBangumiMediaer(), nil
	case global.SourceTMDB:
		return getTMDBMediaer(), nil
	default:
		return nil, ErrNoSuchSource
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
