package service

import (
	"testing"

	"github.com/TravisRoad/goshower/global"
	"github.com/stretchr/testify/assert"
)

func Test_getMediaer(t *testing.T) {
	type args struct {
		source global.Source
		t      global.Type
	}
	tests := []struct {
		name string
		args args
		want Mediaer
	}{
		{"bangumi", args{global.SourceBangumi, global.TypeAnime}, &BangumiMediaer{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := getMediaer(tt.args.source, tt.args.t)
			assert.IsType(t, tt.want, m)
		})
	}
}
