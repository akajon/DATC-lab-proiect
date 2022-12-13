package users

import "time"

type CreateUserResponse struct {
	Id           int
	FirstName    string
	LastName     string
	Email        string
	Password     string
	TaxReduction int
	Role         string
	DeletionDate *time.Time
}

type UpdateDeleteDateResponse struct {
	Id           int
	DeletionDate time.Time
}
