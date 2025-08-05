package user

type Service interface {
	GetUserbyID(id int) (*User, error)
	GetUsers() (*[]User, error)
	PostUser(user *User) error
	PutUser(user *User) error
	DeleteUser(id int) error
}
