package models

import "database/sql"

type User struct {
	Id       int64          `db:"id"`
	UserName sql.NullString `db:"username"`
	Email    sql.NullString `db:"email"`
}
