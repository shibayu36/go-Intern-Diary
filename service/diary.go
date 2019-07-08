package service

import "github.com/hatena/go-Intern-Diary/model"

func (app *diaryApp) CreateNewDiary(userID uint64, name string) error {
	return app.repo.CreateDiary(userID, name)
}

func (app *diaryApp) ListDiariesByUserID(userID uint64) ([]*model.Diary, error) {
	return app.repo.ListDiariesByUserID(userID)
}
