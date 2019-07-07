package service

import (
	"testing"

	"github.com/hatena/go-Intern-Diary/testutil"
	"github.com/stretchr/testify/assert"
)

func TestDiaryApp_CreateNewDiary(t *testing.T) {
	app := newApp()
	defer closeApp(app)

	user := createTestUser(app)

	name := "test diary name " + testutil.RandomString()
	err := app.CreateNewDiary(user.ID, name)

	assert.NoError(t, err)
}
