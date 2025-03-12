package store

import (
	"context"
	"database/sql"
	"time"
)

func (s *UserStore) createUserInvitation(ctx context.Context, tx *sql.Tx, token string, expiry time.Duration, userId int64) error {

	query := `INSERT INTO user_invitations (token, user_id, expiry) VALUES ($1, $2, $3)`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancel()

	_, err := tx.ExecContext(ctx, query, token, userId, time.Now().Add(expiry))
	if err != nil {
		return err
	}

	return nil
}
