package service

import (
	"testing"
	"time"

	"github.com/hatena/go-Intern-Diary/testutil"
	"github.com/stretchr/testify/assert"
)

func TestDiaryApp_CreateNewUser(t *testing.T) {
	app := newApp()
	defer closeApp(app)

	name := "test name " + testutil.RandomString()
	password := testutil.RandomString() + testutil.RandomString()
	err := app.CreateNewUser(name, password)
	assert.NoError(t, err)

	user, err := app.FindUserByName(name)
	assert.NoError(t, err)
	assert.Equal(t, user.Name, name)
}

func TestDiaryApp_CreateNewToken(t *testing.T) {
	app := newApp()
	defer closeApp(app)

	name := "test name " + testutil.RandomString()
	password := testutil.RandomString() + testutil.RandomString()
	err := app.CreateNewUser(name, password)
	assert.NoError(t, err)
	user, _ := app.FindUserByName(name)

	token, err := app.CreateNewToken(user.ID, time.Now().Add(1*time.Hour))
	assert.NoError(t, err)
	assert.NotEqual(t, "", token)

	u, err := app.FindUserByToken(token)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, u.ID)
}

func TestDiaryApp_LoginUser(t *testing.T) {
	app := newApp()
	defer closeApp(app)

	name := "test name " + testutil.RandomString()
	password := testutil.RandomString() + testutil.RandomString()
	err := app.CreateNewUser(name, password)
	assert.NoError(t, err)

	login, err := app.LoginUser(name, password)
	assert.NoError(t, err)
	assert.True(t, login)

	login, err = app.LoginUser(name, password+".")
	assert.NoError(t, err)
	assert.False(t, login)
}
