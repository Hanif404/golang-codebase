package auth

import "golang-codebase/src/domain/user"

type Repository interface {
	FindByID(email string) (*user.User, error)
	Create(registration *Registration) error
}
