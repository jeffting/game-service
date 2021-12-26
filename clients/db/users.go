package db

import (
	"context"
	"database/sql"
	"strings"
)

func (d GameClient) GetUserByUsername(ctx context.Context, username string) (*UserDB, error) {
	// customer := &Entry{}]
	var user UserDB
	// users := []*UserDB{}
	err := d.db.GetContext(
		ctx,
		&user,
		`SELECT *
			FROM users where username = $1`,
		username,
	)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		if strings.Contains(err.Error(), "context canceled") {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}
