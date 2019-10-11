package repository

import (
	"time"

	"github.com/hatena/go-Intern-Diary/model"
	"github.com/jmoiron/sqlx"
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

func (r *repository) ListDiariesByUserIDs(userIDs []uint64) (map[uint64][]*model.Diary, error) {
	if len(userIDs) == 0 {
		return nil, nil
	}
	query, args, err := sqlx.In(
		`SELECT id, user_id, name FROM diary
			WHERE user_id IN (?)`, userIDs,
	)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	diaries := make(map[uint64][]*model.Diary)
	for rows.Next() {
		var diary model.Diary
		rows.Scan(&diary.ID, &diary.UserID, &diary.Name)
		diaries[diary.UserID] = append(diaries[diary.UserID], &diary)
	}
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
