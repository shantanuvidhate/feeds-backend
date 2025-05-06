package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail    = errors.New("a user with this email address already exists")
	ErrDuplicateUsername = errors.New("a user with this username already exists")
)

type User struct {
	ID          int64    `json:"id"`
	Username    string   `json:"username"`
	Email       string   `json:"email"`
	Password    password `json:"-"`
	CreatedAt   string   `json:"created_at"`
	IsActivated bool     `json:"is_activated"`
}

type password struct {
	plainText *string
	hash      []byte
}

type UserStore struct {
	db *sql.DB
}

func (p *password) Set(plainText string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	p.plainText = &plainText
	p.hash = hash
	return nil
}

func (s *UserStore) Create(ctx context.Context, tx *sql.Tx, user *User) error {

	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id, created_at`

	ctx, cancle := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancle()
	err := tx.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Email,
		user.Password.hash,
	).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		case err.Error() == `pq: duplicate key value violates unique constraint "users_username_key"`:
			return ErrDuplicateUsername
		default:
			return err
		}
	}

	return nil
}

func (s *UserStore) GetById(ctx context.Context, userId int64) (*User, error) {
	query := `SELECT id, username, password, email, created_at, is_Activated FROM users WHERE id = $1 AND is_activated = true`

	ctx, cancle := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancle()

	user := &User{}
	err := s.db.QueryRowContext(ctx, query, userId).Scan(
		&user.ID,
		&user.Username,
		&user.Password.hash,
		&user.Email,
		&user.CreatedAt,
		&user.IsActivated,
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

func (s *UserStore) GetByEmail(ctx context.Context, email string) (*User, error) {
	query := `SELECT id, username, password, email, created_at FROM users WHERE email = $1 AND is_activated = true`

	ctx, cancle := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancle()
	user := &User{}
	err := s.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Password.hash,
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

func (s *UserStore) CreateAndInvite(ctx context.Context, user *User, token string, invitationExpiry time.Duration) error {
	return withTx(s.db, ctx, func(tx *sql.Tx) error {
		// Create the user
		if err := s.Create(ctx, tx, user); err != nil {
			return err
		}

		// create the invitation
		if err := s.createUserInvitation(ctx, tx, token, invitationExpiry, user.ID); err != nil {
			return err
		}

		return nil
	})
}

func (s *UserStore) Delete(ctx context.Context, userID int64) error {
	return withTx(s.db, ctx, func(tx *sql.Tx) error {
		if err := s.delete(ctx, tx, userID); err != nil {
			return err
		}

		if err := s.deleteUserInvitations(ctx, tx, userID); err != nil {
			return err
		}

		return nil
	})
}

func (s *UserStore) delete(ctx context.Context, tx *sql.Tx, id int64) error {
	query := `DELETE FROM users WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancel()

	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *password) Compare(plainText string) error {
	return bcrypt.CompareHashAndPassword(p.hash, []byte(plainText))
}
