package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
}

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, user *User) error {

	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id, created_at`

	ctx, cancle := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancle()
	err := s.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Email,
		user.Password,
	).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserStore) GetById(ctx context.Context, userId int64) (*User, error) {
	query := `SELECT id, username, password, email, created_at FROM users WHERE id = $1`

	ctx, cancle := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancle()

	user := &User{}
	err := s.db.QueryRowContext(ctx, query, userId).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
	)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return user, nil

}
