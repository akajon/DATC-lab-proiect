package users

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Repository interface {
	Create(ctx context.Context, firstName, lastName, email, password, role string, taxReduction int, deletionDate *time.Time) (*CreateUserResponse, error)
	UpdateDeleteDate(ctx context.Context, userId int, deleteDate time.Time) (*UpdateDeleteDateResponse, error)
}

func NewService(repo Repository) Service {
	return &serviceImplementation{repo: repo}
}

type serviceImplementation struct {
	repo Repository
}

func (s serviceImplementation) CreateUser(ctx context.Context, firstName, lastName, email, password string) (*CreateUserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, err
	}

	createdUser, err := s.repo.Create(ctx, firstName, lastName, email, string(hashedPassword), "USER", 0, nil)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s serviceImplementation) UpdateDeleteDate(ctx context.Context, userId int) (*UpdateDeleteDateResponse, error) {
	return s.repo.UpdateDeleteDate(ctx, userId, time.Now().AddDate(0, 1, 0))
}
