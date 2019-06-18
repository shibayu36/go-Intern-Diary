package service

func (app *diaryApp) CreateNewDiary(userID uint64, name string) error {
	return app.repo.CreateDiary(userID, name)
}
