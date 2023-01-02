package users

type CreateUserRequest struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type Credentials struct {
	Password string
	Username string
}
