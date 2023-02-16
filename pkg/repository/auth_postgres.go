package repository

import (
	"Brachium/pkg/entity"
	"database/sql"
	"fmt"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db}
}

func (r *AuthPostgres) CreateUser(user entity.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (nickname, email, password) VALUES ($1, $2, $3) RETURNING user_id", usersTable)

	row := r.db.QueryRow(query, user.Nickname, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(email, password string) (entity.User, error) {
	var user entity.User
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE email=$1 AND password=$2", usersTable)
	row := r.db.QueryRow(query, email, password)
	if err := row.Scan(&user.UserId); err != nil {
		return user, err
	}
	return user, nil
}
