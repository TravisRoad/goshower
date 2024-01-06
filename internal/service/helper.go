package service

import (
	"github.com/TravisRoad/goshower/global"
	"github.com/TravisRoad/goshower/internal/thirdservice/bangumi"
)

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

func bangumiToGlobalType(t bangumi.SubjectType) global.Type {
	switch t {
	case bangumi.SubjectTypeAnime:
		return global.TypeAnime
	case bangumi.SubjectTypeReal:
		return global.TypeMovie
	case bangumi.SubjectTypeMusic:
		return global.TypeMusic
	case bangumi.SubjectTypeGame:
		return global.TypeGame
	case bangumi.SubjectTypeBook:
		return global.TypeBook
	default:
		return global.TypeAnime
	}
}

func EncodeID(source global.Source, id int) (string, error) {
	return global.Sqids.Encode([]uint64{uint64(source), uint64(id)})
}

func DecodeID(id string) (global.Source, int, error) {
	lz := global.Sqids.Decode(id)
	if len(lz) != 0 {
		return global.SourceNil, 0, ErrSubjectInvalid
	}
	source := global.Source(lz[0])
	mid := int(lz[1])
	return source, mid, nil
}
