package web

import (
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/hatena/go-Intern-Diary/config"
	"github.com/hatena/go-Intern-Diary/repository"
	"github.com/hatena/go-Intern-Diary/service"
	"github.com/stretchr/testify/assert"
)

func init() {
	csrfMiddleware = func(next http.Handler) http.Handler {
		return next
	}
	csrfToken = func(r *http.Request) string {
		return ""
	}
}

func randomString() string {
	return strconv.FormatInt(time.Now().Unix()^rand.Int63(), 16)
}

func newAppServer() (service.DiaryApp, *httptest.Server) {
	conf, err := config.Load()
	if err != nil {
		panic(err)
	}
	repo, err := repository.New(conf.DbDsn)
	if err != nil {
		panic(err)
	}
	app := service.NewApp(repo)
	handler := NewServer(app).Handler()
	return app, httptest.NewServer(handler)
}

func TestServer_Index(t *testing.T) {
	_, ts := newAppServer()
	defer ts.Close()

	resp, respBody := client.Get(ts.URL + "/").Do()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Contains(t, respBody, `<h1>ダイアリー</h1>`)
}

func TestServer_Signup(t *testing.T) {
	_, ts := newAppServer()
	defer ts.Close()

	resp, respBody := client.Get(ts.URL + "/signup").Do()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Contains(t, respBody, `<h1>ユーザー登録</h1>`)

	name, password := "test name "+randomString(), randomString()
	resp, _ = client.Post(ts.URL+"/signup", map[string]string{
		"name":     name,
		"password": password,
	}).Do()
	location := resp.Header.Get("Location")
	cookie := resp.Cookies()[0]

	assert.Equal(t, http.StatusSeeOther, resp.StatusCode)
	assert.Equal(t, "/", location)
	assert.Equal(t, "DIARY_SESSION", cookie.Name)
	assert.Regexp(t, "^[a-zA-Z0-9_@]{128}$", cookie.Value)

	// TODO: 正しいname, passwordのユーザーが作成されたかを、ログイン処理が出来たら確認
}
