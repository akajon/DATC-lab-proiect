package users

import "time"

type UpdateDeleteDateResponse struct {
	Id           int
	DeletionDate time.Time
}

type SignInResponse struct {
	FirstName    string
	LastName     string
	Email        string
	TaxReduction int
	Role         string
}
