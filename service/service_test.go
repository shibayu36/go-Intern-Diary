package service

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/hatena/go-Intern-Diary/config"
	"github.com/hatena/go-Intern-Diary/repository"
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

// テストでランダム文字列を使いたいときが多い
func randomString() string {
	return strconv.FormatInt(time.Now().Unix()^rand.Int63(), 16)
}
