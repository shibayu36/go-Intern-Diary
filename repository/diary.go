package repository

import (
	"database/sql"
	"time"

	"github.com/hatena/go-Intern-Diary/model"
	"github.com/jmoiron/sqlx"
)

func (r *repository) FindDiaryByID(id uint64) (*model.Diary, error) {
	var diary model.Diary
	err := r.db.Get(
		&diary,
		`SELECT id, user_id, name FROM diary
			WHERE id = ? LIMIT 1`, id,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, userNotFoundError
		}
		return nil, err
	}
	return &diary, nil
}

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

func (r *repository) CreateDiary(userID uint64, name string) (*model.Diary, error) {
	id, err := r.generateID()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	_, err = r.db.Exec(
		`INSERT INTO diary
      (id, user_id, name, created_at, updated_at)
      VALUES (?, ?, ?, ?, ?)`,
		id, userID, name, now, now,
	)
	if err != nil {
		return nil, err
	}

	return r.FindDiaryByID(id)
}
