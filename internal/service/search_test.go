package service

import (
	"testing"

	"github.com/TravisRoad/goshower/global"
	"github.com/stretchr/testify/assert"
)

func TestGetSearcher(t *testing.T) {
	type args struct {
		source global.Source
		t      global.Type
	}
	tests := []struct {
		name string
		args args
		want Searcher
	}{
		{"bangumi", args{global.SourceBangumi, global.TypeAnime}, &BangumiService{}},
		{"tmdb", args{global.SourceTMDB, global.TypeMovie}, &TMDBService{}},
		{"fallback", args{global.SourceNil, global.TypeNil}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			searcher := GetSearcher(tt.args.source, tt.args.t)
			assert.IsType(t, tt.want, searcher)
		})
	}
}
