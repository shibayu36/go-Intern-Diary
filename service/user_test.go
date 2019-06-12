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
}
