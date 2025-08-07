package authrepository

import (
	"database/sql"
	"golang-codebase/src/domain/auth"
	"golang-codebase/src/domain/user"
)

type repo struct {
	db *sql.DB
}

func New(db *sql.DB) auth.Repository {
	return &repo{
		db,
	}
}

func (r *repo) FindByID(email string) (*user.User, error) {
	var user user.User

	// Execute the query
	result := r.db.QueryRow("SELECT id, name, email, password FROM user where email = ?", email)
	err := result.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // not found
		}
		return nil, err
	}
	return &user, nil
}

func (r *repo) Create(user *auth.Registration) error {
	_, err := r.db.Exec("INSERT INTO user (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)
	return err
}
