package web

import (
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/hatena/go-Intern-Diary/config"
	"github.com/hatena/go-Intern-Diary/model"
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

func createTestUser(app service.DiaryApp) *model.User {
	name := "test name " + randomString()
	password := randomString() + randomString()
	err := app.CreateNewUser(name, password)
	if err != nil {
		panic(err)
	}
	user, err := app.FindUserByName(name)
	if err != nil {
		panic(err)
	}
	return user
}

func TestServer_Index(t *testing.T) {
	_, ts := newAppServer()
	defer ts.Close()

	resp, respBody := client.Get(ts.URL + "/").Do()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Contains(t, respBody, `<h1>ダイアリー</h1>`)
}
