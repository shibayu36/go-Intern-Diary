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

func TestDiaryApp_ListDiariesByUserID(t *testing.T) {
	app := newApp()
	defer closeApp(app)

	user := createTestUser(app)

	name1 := "test diary1 " + testutil.RandomString()
	app.CreateNewDiary(user.ID, name1)

	// 最初は一つだけダイアリーを持つ
	diaries, err := app.ListDiariesByUserID(user.ID)
	assert.NoError(t, err)
	assert.Len(t, diaries, 1)
	assert.Equal(t, name1, diaries[0].Name)

	name2 := "test diary2 " + testutil.RandomString()
	app.CreateNewDiary(user.ID, name2)

	// 二つになった
	diaries, err = app.ListDiariesByUserID(user.ID)
	assert.NoError(t, err)
	assert.Len(t, diaries, 2)
	assert.Equal(t, name2, diaries[0].Name)
	assert.Equal(t, name1, diaries[1].Name)
}
