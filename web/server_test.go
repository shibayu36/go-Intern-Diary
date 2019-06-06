package web

import (
	"net/http"
	"net/http/httptest"
	"testing"

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
}
