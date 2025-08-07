package authservice

import (
	"golang-codebase/src/domain/auth"
	"golang-codebase/src/utils"

	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo auth.Repository
}

func New(r auth.Repository) auth.Service {
	return &service{
		repo: r,
	}
}

func (s *service) Login(login *auth.Login) (string, error) {
	result, err := s.repo.FindByID(login.Email)
	if err != nil {
		return "", err
	}

	//compare password
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(login.Password))
	if err != nil {
		return "", err
	}

	//generate jwt
	t, err := utils.GenerateToken(uint(result.ID))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (s *service) Registration(reg *auth.Registration) error {
	//hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reg.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	reg.Password = string(hashedPassword)
	return s.repo.Create(reg)
}
