package service

import (
	"github.com/TravisRoad/goshower/global"
	"github.com/TravisRoad/goshower/internal/model"
)

type ISearchService interface {
	Search(query string, page, size int, st SourceType) (*model.SearchResult, error)
}

// MediaService
type IMediaService interface {
	AddMedia(id int, source global.Source, t global.Type) error
	MediaDetail(id int, source global.Source, t global.Type) (*model.Media, error)
}

type Searcher interface {
	Search(query string, page, size int) (*model.SearchResult, error)
}

type Mediaer interface {
	MediaDetail(id int) (*model.Media, error)
	AddMedia(id int) error
}
