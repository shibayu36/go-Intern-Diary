package service

import (
	"errors"

	"github.com/hatena/go-Intern-Diary/model"
	"golang.org/x/crypto/bcrypt"
)

// ユーザ名とパスワードからユーザーを作成
func (app *diaryApp) CreateNewUser(name string, password string) (err error) {
	if name == "" {
		return errors.New("empty user name")
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return app.repo.CreateNewUser(name, string(passwordHash))
}

// ユーザー名からユーザーを取得
func (app *diaryApp) FindUserByName(name string) (*model.User, error) {
	return app.repo.FindUserByName(name)
}
