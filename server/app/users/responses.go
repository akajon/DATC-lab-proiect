package users

import "time"

type UpdateDeleteDateResponse struct {
	Id           int
	DeletionDate time.Time
}
