package service

import (
	"math/rand"
	"time"

	"github.com/hatena/go-Intern-Diary/model"
	"github.com/hatena/go-Intern-Diary/repository"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type DiaryApp interface {
	Close() error

	CreateNewUser(name string, passwordHash string) error
	FindUserByName(name string) (*model.User, error)
	CreateNewToken(userID uint64, expiresAt time.Time) (string, error)
	LoginUser(name string, password string) (bool, error)
	FindUserByToken(token string) (*model.User, error)

	CreateNewDiary(userID uint64, name string) error
	ListDiariesByUserID(userID uint64) ([]*model.Diary, error)
}

func NewApp(repo repository.Repository) DiaryApp {
	return &diaryApp{repo: repo}
}

type diaryApp struct {
	repo repository.Repository
}

func (app *diaryApp) Close() error {
	return app.repo.Close()
}
