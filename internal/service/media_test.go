package service

import (
	"testing"

	"github.com/TravisRoad/goshower/global"
	"github.com/stretchr/testify/assert"
)

func Test_getMediaer(t *testing.T) {
	type args struct {
		source global.Source
	}
	tests := []struct {
		name string
		args args
		want Mediaer
	}{
		{"bangumi", args{global.SourceBangumi}, &BangumiService{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := getMediaer(tt.args.source)
			assert.IsType(t, tt.want, m)
		})
	}
}
