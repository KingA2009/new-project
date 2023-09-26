package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/zhashkevych/todo-app"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf(AppendUsersQuery)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password, user.Email, user.Telephone, user.Address)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (todo.User, error) {
	var user todo.User
	query := fmt.Sprintf(GetUserQuery)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
