package service

import (
	"github.com/TravisRoad/goshower/global"
	"github.com/TravisRoad/goshower/internal/model"
)

type ISearchService interface {
	Search(query string, page, size int, source global.Source, t global.Type) (*model.SearchResult, error)
}

// MediaService
type IMediaService interface {
	MediaDetail(id int, source global.Source) (*model.Media, error)
}

type Searcher interface {
	Search(query string, page, size int) (*model.SearchResult, error)
}

type Mediaer interface {
	MediaDetail(id int) (*model.Media, error)
}

type Recorder interface {
	RecordSubject(id int, src global.Source, uid uint, action global.Status) error
	RecordEp(id int, src global.Source, uid uint, action global.Status, ep int) error
}
