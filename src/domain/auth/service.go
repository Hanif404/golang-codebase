package auth

type Service interface {
	Login(login *Login) (string, error)
	Registration(registration *Registration) error
}
