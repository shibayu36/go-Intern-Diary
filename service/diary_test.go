package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiaryApp_CreateNewDiary(t *testing.T) {
	app := newApp()
	defer closeApp(app)

	user := createTestUser(app)

	name := "test diary name " + randomString()
	err := app.CreateNewDiary(user.ID, name)

	assert.NoError(t, err)
}
