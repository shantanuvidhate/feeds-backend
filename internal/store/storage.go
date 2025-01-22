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
		Delete(context.Context, int64) error
		Update(context.Context, *Post) error
	}
	User interface {
		Create(context.Context, *User) error
	}
	Comment interface {
		GetByPostId(context.Context, int64) ([]Comment, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Post:    &PostStore{db},
		User:    &UserStore{db},
		Comment: &CommentStore{db},
	}
}
