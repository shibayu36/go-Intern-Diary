package repository

import (
	"time"
)

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
