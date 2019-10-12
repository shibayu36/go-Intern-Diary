package service

import "github.com/hatena/go-Intern-Diary/model"

func (app *diaryApp) CreateNewDiary(userID uint64, name string) error {
	return app.repo.CreateDiary(userID, name)
}

func (app *diaryApp) FindDiaryByID(id uint64) (*model.Diary, error) {
	return app.repo.FindDiaryByID(id)
}

func (app *diaryApp) ListDiariesByUserID(userID uint64) ([]*model.Diary, error) {
	return app.repo.ListDiariesByUserID(userID)
}

func (app *diaryApp) ListDiariesByUserIDs(userIDs []uint64) (map[uint64][]*model.Diary, error) {
	return app.repo.ListDiariesByUserIDs(userIDs)
}
