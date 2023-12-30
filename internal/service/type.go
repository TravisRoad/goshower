package service

import (
	"github.com/TravisRoad/goshower/global"
	"github.com/TravisRoad/goshower/internal/thirdservice/bangumi"
)

type SourceType struct {
	Source global.Source
	Type   global.Type
}

func toBangumiType(t global.Type) bangumi.SubjectType {
	switch t {
	case global.TypeAnime:
		return bangumi.SubjectTypeAnime
	case global.TypeMovie:
		return bangumi.SubjectTypeReal
	case global.TypeMusic:
		return bangumi.SubjectTypeMusic
	case global.TypeGame:
		return bangumi.SubjectTypeGame
	case global.TypeBook:
		return bangumi.SubjectTypeBook
	default:
		return bangumi.SubjectTypeAnime
	}
}
