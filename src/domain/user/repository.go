package user

type Repository interface {
	FindByID(id int) (*User, error)
	FindAll() (*[]User, error)
	Save(*User) error
	Update(*User) error
	Delete(id int) error
}
