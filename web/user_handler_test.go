package web

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestServer_Signup(t *testing.T) {
	app, ts := newAppServer()
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

	// 作ったユーザーでログインできる
	loginSuccess, _ := app.LoginUser(name, password)
	assert.Equal(t, true, loginSuccess)
}

func TestServer_Signin(t *testing.T) {
	app, testServer := newAppServer()
	defer testServer.Close()

	resp, respBody := client.Get(testServer.URL + "/signin").Do()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Contains(t, respBody, `<h1>ログイン</h1>`)

	name, password := "test name "+randomString(), randomString()
	err := app.CreateNewUser(name, password)
	assert.NoError(t, err)
	resp, _ = client.Post(testServer.URL+"/signin", map[string]string{
		"name":     name,
		"password": password,
	}).Do()
	location := resp.Header.Get("Location")
	cookie := resp.Cookies()[0]

	assert.Equal(t, http.StatusSeeOther, resp.StatusCode)
	assert.Equal(t, "/", location)
	assert.Equal(t, "DIARY_SESSION", cookie.Name)
	assert.Regexp(t, "^[a-zA-Z0-9_@]{128}$", cookie.Value)
}

func TestServer_Signout(t *testing.T) {
	app, testServer := newAppServer()
	defer testServer.Close()

	user := createTestUser(app)
	expiresAt := time.Now().Add(24 * time.Hour)
	token, _ := app.CreateNewToken(user.ID, expiresAt)
	sessionCookie := &http.Cookie{Name: sessionKey, Value: token, Expires: expiresAt}

	resp, respBody := client.Get(testServer.URL + "/").WithCookie(sessionCookie).Do()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Contains(t, respBody, "ユーザー名: "+user.Name)
	assert.Contains(t, respBody, `<input type="submit" value="ログアウト"/>`)

	resp, _ = client.Post(testServer.URL+"/signout", nil).WithCookie(sessionCookie).Do()
	location := resp.Header.Get("Location")

	assert.Equal(t, http.StatusSeeOther, resp.StatusCode)
	assert.Equal(t, "/", location)
	var cookie *http.Cookie
	for _, c := range resp.Cookies() {
		if c.Name == sessionKey {
			cookie = c
		}
	}
	assert.Equal(t, "", cookie.Value)
}
