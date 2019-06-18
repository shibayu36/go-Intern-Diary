package model

type Diary struct {
	ID     uint64 `db:"id"`
	UserID uint64 `db:"user_id"`
	Name   string `db:"name"`
}
