package service

import (
	"github.com/hatena/go-Intern-Diary/config"
	"github.com/hatena/go-Intern-Diary/model"
	"github.com/hatena/go-Intern-Diary/repository"
	"github.com/hatena/go-Intern-Diary/testutil"
)

// テスト用にDiaryAppを作るユーティリティ
func newApp() DiaryApp {
	conf, err := config.Load()
	if err != nil {
		panic(err)
	}
	repo, err := repository.New(conf.DbDsn)
	if err != nil {
		panic(err)
	}
	return NewApp(repo)
}

func closeApp(app DiaryApp) {
	err := app.Close()
	if err != nil {
		panic(err)
	}
}

// テスト用ユーザーの作成
func createTestUser(app DiaryApp) *model.User {
	name := "test name " + testutil.RandomString()
	password := testutil.RandomString() + testutil.RandomString()
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
