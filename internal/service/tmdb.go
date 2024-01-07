package service

import (
	"fmt"
	"time"

	"github.com/TravisRoad/goshower/global"
	"github.com/TravisRoad/goshower/internal/model"
	"github.com/TravisRoad/goshower/internal/thirdservice/tmdb/v3"
)

type TMDBService struct {
	Source global.Source
	Client *tmdb.Client
}

func (s *TMDBService) MediaDetail(id int) (*model.Media, error) {
	detail, err := s.Client.MovieDetail(id, tmdb.MovieDetailOption{Language: "zh-CN"})
	if err != nil {
		return nil, err
	}
	pubDate, err := time.Parse("2006-01-02", detail.ReleaseDate)
	if err != nil {
		return nil, err
	}
	res := &model.Media{
		Source:      s.Source,
		Type:        global.TypeMovie, // FIXME: I assert this is the movie type
		Link:        fmt.Sprintf("%s/%d", tmdb.MovieURL, id),
		MediaID:     detail.ID,
		Title:       detail.OriginalTitle,
		TitleCn:     detail.Title,
		Summary:     detail.Overview,
		PublishData: pubDate,
		Nsfw:        detail.Adult,
		Platform:    "",
		ImageLarge:  fmt.Sprintf("%s/%s/%s", tmdb.PicURL, "original", detail.PosterPath), // https://image.tmdb.org/t/p/w1280/6pZgH10jhpToPcf0uvyTCPFhWpI.jpg
		ImageCommon: fmt.Sprintf("%s/%s/%s", tmdb.PicURL, "w780", detail.PosterPath),     // https://image.tmdb.org/t/p/w1280/6pZgH10jhpToPcf0uvyTCPFhWpI.jpg
		ImageMedium: fmt.Sprintf("%s/%s/%s", tmdb.PicURL, "w500", detail.PosterPath),     // https://image.tmdb.org/t/p/w1280/6pZgH10jhpToPcf0uvyTCPFhWpI.jpg
		Eps:         1,
		RatingScore: float32(detail.VoteAverage),
	}

	return res, nil
}

func (s *TMDBService) Search(query string, page, size int) (*model.SearchResult, error) {
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
		if len(item.PosterPath) != 0 {
			searchItem.Pic = fmt.Sprintf("%s/%s/%s", tmdb.PicURL, "w500", item.PosterPath)
		}
		id, err := global.Sqids.Encode([]uint64{uint64(s.Source), uint64(item.ID)})
		if err != nil {
			return nil, err
		}
		searchItem.ID = id
		sr.Items[i] = searchItem
	}
	return sr, nil
}
