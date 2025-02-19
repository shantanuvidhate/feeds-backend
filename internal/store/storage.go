package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrRecordNotFound    = errors.New("record not found")
	ErrConflict          = errors.New("record already exists")
	QueryTimeOutDuration = 5 * time.Second
)

type Storage struct {
	Post interface {
		Create(context.Context, *Post) error
		GetById(context.Context, int64) (*Post, error)
		Delete(context.Context, int64) error
		Update(context.Context, *Post) error
	}
	User interface {
		GetById(context.Context, int64) (*User, error)
		Create(context.Context, *User) error
	}
	Comment interface {
		Create(context.Context, *Comment) error
		GetByPostId(context.Context, int64) ([]Comment, error)
	}
	Followers interface {
		Follow(ctx context.Context, followerId int64, userId int64) error
		Unfollow(ctx context.Context, followerId int64, userId int64) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Post:      &PostStore{db},
		User:      &UserStore{db},
		Comment:   &CommentStore{db},
		Followers: &FollowerStore{db},
	}
}
