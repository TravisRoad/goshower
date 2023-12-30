package tmdb

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (fn RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return fn(req), nil
}

func TestClient_SearchMovie(t *testing.T) {
	_t := func(t *testing.T, query string, opt SearchMovieOption, rtfn RoundTripFunc, wantErr bool) {
		cli := &Client{
			Cli: &http.Client{
				Transport: rtfn,
			},
			Token: "",
		}
		_, err := cli.SearchMovie(query, opt)
		if wantErr {
			assert.NotNil(t, err)
		}
	}

	var (
		respjson = []string{
			"{\"page\":1,\"results\":[{\"adult\":false,\"backdrop_path\":\"/tjMiLkVfOmbx3kUtKSmkLimiw8x.jpg\",\"genre_ids\":[14,16,12],\"id\":4935,\"original_language\":\"ja\",\"original_title\":\"ハウルの動く城\",\"overview\":\"继母因无力负担生活，将苏菲和她的两个妹妹都送到了制帽店去当学徒。两个妹妹很快先后就离开了制帽店去追寻各自的梦想，只有苏菲坚持了下来。一天，小镇旁边来了一座移动堡垒，传说堡垒的主人哈尔专吸取年青姑娘的灵魂，所以小镇的姑娘都不敢靠近。一个恶毒的巫婆嫉妒苏菲的制帽技术，用巫术把她变成了一个80岁的老太婆，而且苏菲还不能对别人说出自己身中的巫术。无奈，苏菲决定独自一人逃离小镇。天黑了，虚弱的苏菲没走多远，来到了移动城堡。心想自己已经是老太婆了，苏菲壮着胆子走进了城堡。不想，遇到了和她遭遇相同的火焰魔。两人约定彼此帮助对方打破各自的咒语。\",\"popularity\":106.491,\"poster_path\":\"/tuOu8C02KULf75hehYS6Eowen4a.jpg\",\"release_date\":\"2004-09-09\",\"title\":\"哈尔的移动城堡\",\"video\":false,\"vote_average\":8.405,\"vote_count\":8983}],\"total_pages\":1,\"total_results\":1}",
			"{\"page\":\"1\",\"results\":[{\"adult\":false,\"backdrop_path\":\"/tjMiLkVfOmbx3kUtKSmkLimiw8x.jpg\",\"genre_ids\":[14,16,12],\"id\":4935,\"original_language\":\"ja\",\"original_title\":\"ハウルの動く城\",\"overview\":\"继母因无力负担生活，将苏菲和她的两个妹妹都送到了制帽店去当学徒。两个妹妹很快先后就离开了制帽店去追寻各自的梦想，只有苏菲坚持了下来。一天，小镇旁边来了一座移动堡垒，传说堡垒的主人哈尔专吸取年青姑娘的灵魂，所以小镇的姑娘都不敢靠近。一个恶毒的巫婆嫉妒苏菲的制帽技术，用巫术把她变成了一个80岁的老太婆，而且苏菲还不能对别人说出自己身中的巫术。无奈，苏菲决定独自一人逃离小镇。天黑了，虚弱的苏菲没走多远，来到了移动城堡。心想自己已经是老太婆了，苏菲壮着胆子走进了城堡。不想，遇到了和她遭遇相同的火焰魔。两人约定彼此帮助对方打破各自的咒语。\",\"popularity\":106.491,\"poster_path\":\"/tuOu8C02KULf75hehYS6Eowen4a.jpg\",\"release_date\":\"2004-09-09\",\"title\":\"哈尔的移动城堡\",\"video\":false,\"vote_average\":8.405,\"vote_count\":8983}],\"total_pages\":1,\"total_results\":1}",
		}
	)

	tests := []struct {
		name    string
		query   string
		opt     SearchMovieOption
		rtfn    RoundTripFunc
		wantErr bool
	}{
		{
			name:  "success",
			query: "哈尔的移动",
			opt: SearchMovieOption{
				Adult:    true,
				Language: "zh-CN",
				Page:     1,
				Region:   "",
				Year:     0,
			},
			rtfn: RoundTripFunc(func(req *http.Request) *http.Response {
				param := req.URL.Query()
				assert.Equal(t, "/3/search/movie", req.URL.Path)
				assert.Equal(t, "哈尔的移动", param.Get("query"))
				assert.Equal(t, "true", param.Get("include_adult"))
				assert.Equal(t, "zh-CN", param.Get("language"))
				assert.Equal(t, "1", param.Get("page"))
				assert.Len(t, param.Get("region"), 0)
				assert.Len(t, param.Get("year"), 0)
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(respjson[0])),
				}
			}),
			wantErr: false,
		},
		{
			name:  "json decode error",
			query: "哈尔的移动",
			opt:   SearchMovieOption{},
			rtfn: RoundTripFunc(func(req *http.Request) *http.Response {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(respjson[1])),
				}
			}),
			wantErr: true,
		},
		{
			name:  "retuan status not ok",
			query: "query",
			opt:   SearchMovieOption{},
			rtfn: RoundTripFunc(func(req *http.Request) *http.Response {
				return &http.Response{
					StatusCode: http.StatusUnauthorized,
					Body:       io.NopCloser(strings.NewReader(respjson[0])),
				}
			}),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_t(t, tt.query, tt.opt, tt.rtfn, tt.wantErr)
		})
	}
}
