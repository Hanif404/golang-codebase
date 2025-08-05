package userservice

import "golang-codebase/src/domain/user"

type service struct {
	repo user.Repository
}

func New(r user.Repository) user.Service {
	return &service{
		repo: r,
	}
}

func (s *service) GetUserbyID(id int) (*user.User, error) {
	return s.repo.FindByID(id)
}

func (s *service) GetUsers() (*[]user.User, error) {
	return s.repo.FindAll()
}

func (s *service) PostUser(user *user.User) error {
	return s.repo.Save(user)
}

func (s *service) PutUser(user *user.User) error {
	return s.repo.Update(user)
}

func (s *service) DeleteUser(id int) error {
	return s.repo.Delete(id)
}
