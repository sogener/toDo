package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"toDo"
)

type AuthPostgres struct {
	db *sqlx.DB
}

// NewAuthPostgres this is constructor
func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user toDo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO \"%s\" (name, username, password) values ($1, $2, $3) RETURNING id", userTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
