package models

import "database/sql"

type User struct {
	ID int      `json:"id"`
	Name string `json:"name"`
	Age int     `json:"age"`
	Email sql.NullString `json:"email"`
}
