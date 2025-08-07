package auth

type Login struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type Registration struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}
