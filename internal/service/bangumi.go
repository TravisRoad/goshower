package service

import (
	"fmt"
	"math"
	"time"

	"github.com/TravisRoad/goshower/global"
	"github.com/TravisRoad/goshower/internal/model"
	"github.com/TravisRoad/goshower/internal/thirdservice/bangumi"
)

type BangumiService struct {
	Source global.Source
	Type   bangumi.SubjectType
	Client *bangumi.Client
}

func (s *BangumiService) MediaDetail(id int) (*model.Media, error) {
	detail, err := s.Client.GetSubjectDetail(id)
	if err != nil {
		return nil, err
	}
	pubDate, err := time.Parse("2006-01-02", detail.Date)
	if err != nil {
		return nil, err
	}
	res := &model.Media{
		Source:      s.Source,
		Type:        bangumiToGlobalType(bangumi.SubjectType(detail.Type)),
		Link:        fmt.Sprintf("%s/%d", bangumi.SubjectURL, detail.ID),
		MediaID:     detail.ID,
		Title:       detail.Name,
		TitleCn:     detail.NameCn,
		Summary:     detail.Summary,
		PublishData: pubDate,
		Status:      0,
		StatusText:  "",
		Nsfw:        detail.Nsfw,
		Platform:    detail.Platform,
		ImageLarge:  detail.Images.Large,
		ImageCommon: detail.Images.Common,
		ImageMedium: detail.Images.Medium,
		Eps:         detail.Eps,
		RatingScore: float32(detail.Rating.Score),
	}
	return res, nil
}

func (s *BangumiService) Search(query string, page, size int) (*model.SearchResult, error) {
	maxResult := min(size, 25)
	start := min((page-1)*maxResult, 0)

	searchResult, err := s.Client.Search(query, bangumi.SearchOption{
		Start:         start,
		MaxResult:     maxResult,
		Type:          int(s.Type),
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

		id, err := EncodeID(s.Source, item.ID)
		if err != nil {
			return nil, err
		}

		searchItem.ID = id
		sr.Items = append(sr.Items, searchItem)
	}

	return sr, nil
}
