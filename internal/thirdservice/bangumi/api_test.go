package bangumi

import (
	"fmt"
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

func TestClient_Search(t *testing.T) {
	respJson := "{\"results\":23,\"list\":[{\"id\":2682,\"url\":\"http://bgm.tv/subject/2682\",\"type\":2,\"name\":\"ぷちえゔぁ\",\"name_cn\":\"EVA爆笑学园\",\"summary\":\"\",\"air_date\":\"\",\"air_weekday\":0,\"images\":{\"large\":\"http://lain.bgm.tv/pic/cover/l/f1/5e/2682_6162O.jpg\",\"common\":\"http://lain.bgm.tv/pic/cover/c/f1/5e/2682_6162O.jpg\",\"medium\":\"http://lain.bgm.tv/pic/cover/m/f1/5e/2682_6162O.jpg\",\"small\":\"http://lain.bgm.tv/pic/cover/s/f1/5e/2682_6162O.jpg\",\"grid\":\"http://lain.bgm.tv/pic/cover/g/f1/5e/2682_6162O.jpg\"}}]}"

	opt := SearchOption{
		Start:         1,
		MaxResult:     25,
		Type:          5,
		ResponseGroup: "large",
	}
	query := "test"

	cli := &Client{
		Cli: &http.Client{
			Transport: RoundTripFunc(func(req *http.Request) *http.Response {
				param := req.URL.Query()
				assert.Equal(t, fmt.Sprintf("/search/subject/%s", query), req.URL.Path)
				assert.Equal(t, fmt.Sprint(opt.Start), param.Get("start"))
				assert.Equal(t, fmt.Sprint(opt.MaxResult), param.Get("max_results"))
				assert.Equal(t, fmt.Sprint(opt.Type), param.Get("type"))
				assert.Equal(t, opt.ResponseGroup, param.Get("responseGroup"))

				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(respJson)),
				}
			}),
		},
		Token:     "token",
		UserAgent: "useragent",
	}
	_, err := cli.Search(query, opt)
	assert.Nil(t, err)
}

func TestClient_GetSubjectDetail(t *testing.T) {
	sid := 42
	respJson := "{\"id\":0,\"type\":2,\"name\":\"string\",\"name_cn\":\"string\",\"summary\":\"string\",\"nsfw\":true,\"locked\":true,\"date\":\"string\",\"platform\":\"string\",\"images\":{\"large\":\"string\",\"common\":\"string\",\"medium\":\"string\",\"small\":\"string\",\"grid\":\"string\"},\"infobox\":[{\"key\":\"简体中文名\",\"value\":\"鲁路修·兰佩路基\"},{\"key\":\"别名\",\"value\":[{\"v\":\"L.L.\"},{\"v\":\"勒鲁什\"},{\"v\":\"鲁鲁修\"},{\"v\":\"ゼロ\"},{\"v\":\"Zero\"},{\"k\":\"英文名\",\"v\":\"Lelouch Lamperouge\"},{\"k\":\"第二中文名\",\"v\":\"鲁路修·冯·布里塔尼亚\"},{\"k\":\"英文名二\",\"v\":\"Lelouch Vie Britannia\"},{\"k\":\"日文名\",\"v\":\"ルルーシュ・ヴィ・ブリタニア\"}]},{\"key\":\"性别\",\"value\":\"男\"},{\"key\":\"生日\",\"value\":\"12月5日\"},{\"key\":\"血型\",\"value\":\"A型\"},{\"key\":\"身高\",\"value\":\"178cm→181cm\"},{\"key\":\"体重\",\"value\":\"54kg\"},{\"key\":\"引用来源\",\"value\":\"Wikipedia\"}],\"volumes\":0,\"eps\":0,\"total_episodes\":0,\"rating\":{\"rank\":0,\"total\":0,\"count\":{\"1\":0,\"2\":0,\"3\":0,\"4\":0,\"5\":0,\"6\":0,\"7\":0,\"8\":0,\"9\":0,\"10\":0},\"score\":0},\"collection\":{\"wish\":0,\"collect\":0,\"doing\":0,\"on_hold\":0,\"dropped\":0},\"tags\":[{\"name\":\"string\",\"count\":0}]}"
	cli := &Client{
		Cli: &http.Client{
			Transport: RoundTripFunc(func(req *http.Request) *http.Response {
				assert.Equal(t, fmt.Sprintf("/v0/subjects/%d", sid), req.URL.Path)

				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(respJson)),
				}
			}),
		},
		Token:     "token",
		UserAgent: "useragent",
	}
	_, err := cli.GetSubjectDetail(sid)
	assert.Nil(t, err)
}
