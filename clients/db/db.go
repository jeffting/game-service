package db

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Game interface {
	GetUserByUsername(ctx context.Context, username string) (*UserDB, error)
	// GetContext(context.Context, interface{}, string, ...interface{}) error
	// SelectContext(context.Context, interface{}, string, ...interface{}) error
}

type GameClient struct {
	db *sqlx.DB
}

func NewGameClient() Game {
	// dbPassword := "gamingdb"
	// dbEndpoint := "localhost:5432"
	// userName := "userjeff"

	// connStr := fmt.Sprintf("postgres://%s:%s@%s",
	// 	userName,
	// 	url.QueryEscape(dbPassword),
	// 	dbEndpoint,
	// )
	connStr := "postgres://userjeff:gamingdb@localhost:5432/postgres?sslmode=disable"

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(300 * time.Second)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	client := GameClient{
		db: db,
	}
	return &client
}
