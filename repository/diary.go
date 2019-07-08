package repository

import (
	"time"

	"github.com/hatena/go-Intern-Diary/model"
)

func (r *repository) ListDiariesByUserID(userID uint64) ([]*model.Diary, error) {
	diaries := []*model.Diary{}
	err := r.db.Select(
		&diaries,
		`SELECT id, user_id, name FROM diary
			WHERE user_id = ?
			ORDER BY created_at DESC`,
		userID,
	)
	return diaries, err
}

func (r *repository) CreateDiary(userID uint64, name string) error {
	id, err := r.generateID()
	if err != nil {
		return err
	}

	now := time.Now()
	_, err = r.db.Exec(
		`INSERT INTO diary
      (id, user_id, name, created_at, updated_at)
      VALUES (?, ?, ?, ?, ?)`,
		id, userID, name, now, now,
	)
	return err
}
