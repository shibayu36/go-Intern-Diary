package service

import (
	"errors"

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
