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
		name    string
		args    args
		want    Mediaer
		wantErr error
	}{
		{global.SourceBangumi.String(), args{global.SourceBangumi}, &BangumiService{}, nil},
		{global.SourceTMDB.String(), args{global.SourceTMDB}, &TMDBService{}, nil},
		{"other", args{}, nil, ErrNoSuchSource},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := getMediaer(tt.args.source)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.IsType(t, tt.want, m)
		})
	}
}
