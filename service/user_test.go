package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDiaryApp_CreateNewUser(t *testing.T) {
	app := newApp()
	defer closeApp(app)

	name := "test name " + randomString()
	password := randomString() + randomString()
	err := app.CreateNewUser(name, password)
	assert.NoError(t, err)

	user, err := app.FindUserByName(name)
	assert.NoError(t, err)
	assert.Equal(t, user.Name, name)
}

func TestDiaryApp_CreateNewToken(t *testing.T) {
	app := newApp()
	defer closeApp(app)

	name := "test name " + randomString()
	password := randomString() + randomString()
	err := app.CreateNewUser(name, password)
	assert.NoError(t, err)
	user, _ := app.FindUserByName(name)

	token, err := app.CreateNewToken(user.ID, time.Now().Add(1*time.Hour))
	assert.NoError(t, err)
	assert.NotEqual(t, "", token)

	// TODO: ユーザをトークンで取得できるようになったらテスト追加
}

func TestDiaryApp_LoginUser(t *testing.T) {
	app := newApp()
	defer closeApp(app)

	name := "test name " + randomString()
	password := randomString() + randomString()
	err := app.CreateNewUser(name, password)
	assert.NoError(t, err)

	login, err := app.LoginUser(name, password)
	assert.NoError(t, err)
	assert.True(t, login)

	login, err = app.LoginUser(name, password+".")
	assert.NoError(t, err)
	assert.False(t, login)
}
