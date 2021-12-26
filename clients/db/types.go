package db

import "time"

type UserDB struct {
	ID        string     `db:"id"`
	Username  string     `db:"username"`
	Password  string     `db:"password"`
	Email     string     `db:"email"`
	CreatedOn *time.Time `db:"created_on"`
	LastLogin *time.Time `db:"last_login"`
}
