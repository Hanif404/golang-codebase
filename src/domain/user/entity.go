package user

type User struct {
	ID       int
	Name     string `validate:"required,min=3"`
	Email    string `validate:"required,email"`
	Password string
}
