package store

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Storage struct {
	Post interface {
		Create(context.Context, *Post) error
		GetById(context.Context, int64) (*Post, error)
	}
	User interface {
		Create(context.Context, *User) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Post: &PostStore{db},
		User: &UserStore{db},
	}
}
