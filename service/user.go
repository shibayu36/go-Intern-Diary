package service

import (
	"crypto/rand"
	"errors"
	"time"

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

// ユーザーIDからユーザーを取得
func (app *diaryApp) FindUserByID(id uint64) (*model.User, error) {
	return app.repo.FindUserByID(id)
}

// ユーザー名からユーザーを取得
func (app *diaryApp) FindUserByName(name string) (*model.User, error) {
	return app.repo.FindUserByName(name)
}

// ユーザーのログイン処理
func (app *diaryApp) LoginUser(name string, password string) (bool, error) {
	passwordHash, err := app.repo.FindPasswordHashByName(name)
	if err != nil {
		return false, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// ログイン用トークンの発行
func (app *diaryApp) CreateNewToken(userID uint64, expiresAt time.Time) (string, error) {
	token := generateToken()
	err := app.repo.CreateNewToken(userID, token, expiresAt)
	if err != nil {
		return "", err
	}
	return token, nil
}

func generateToken() string {
	table := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_@"
	l := len(table)
	ret := make([]byte, 128)
	src := make([]byte, 128)
	rand.Read(src)
	for i := 0; i < 128; i++ {
		ret[i] = table[int(src[i])%l]
	}
	return string(ret)
}

func (app *diaryApp) FindUserByToken(token string) (*model.User, error) {
	return app.repo.FindUserByToken(token)
}
