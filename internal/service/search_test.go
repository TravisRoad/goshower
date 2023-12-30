package service

import (
	"testing"

	"github.com/TravisRoad/goshower/global"
	"github.com/stretchr/testify/assert"
)

func TestGetSearcher(t *testing.T) {
	type args struct {
		st SourceType
	}
	tests := []struct {
		name string
		args args
		want Searcher
	}{
		{"bangumi", args{SourceType{global.SourceBangumi, global.TypeAnime}}, &BangumiSearchService{}},
		{"fallback", args{SourceType{global.SourceOther, global.TypeAnime}}, &FallBackSearchService{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			searcher := GetSearcher(tt.args.st)
			assert.IsType(t, tt.want, searcher)
		})
	}
}
