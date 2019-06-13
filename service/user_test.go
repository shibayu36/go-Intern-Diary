package service

import (
	"testing"

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
