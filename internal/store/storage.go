package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Post interface {
		Create(context.Context, *Post) error
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
