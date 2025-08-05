package userrepository

import (
	"database/sql"
	"golang-codebase/src/domain/user"
)

type repo struct {
	db *sql.DB
}

func New(db *sql.DB) user.Repository {
	return &repo{
		db,
	}
}

func (r *repo) FindByID(id int) (*user.User, error) {
	var user user.User

	// Execute the query
	result := r.db.QueryRow("SELECT id, name, email FROM user where ID = ?", id)
	err := result.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // not found
		}
		return nil, err
	}
	return &user, nil
}

func (r *repo) FindAll() (*[]user.User, error) {
	// Execute the query
	results, err := r.db.Query("SELECT id, name, email FROM user")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	var userList []user.User
	for results.Next() {
		var user user.User
		err = results.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil // not found
			}
			return nil, err
		}
		userList = append(userList, user)
	}

	return &userList, nil
}

func (r *repo) Save(user *user.User) error {
	_, err := r.db.Exec("INSERT INTO user (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) Update(user *user.User) error {
	_, err := r.db.Exec("UPDATE user SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM user WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
