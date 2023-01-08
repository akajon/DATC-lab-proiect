package users

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	Create(ctx context.Context, firstName, lastName, email, password, role string, taxReduction int, deletionDate *time.Time) error
	UpdateDeleteDate(ctx context.Context, userId int, deleteDate time.Time) error
	PasswordAndId(ctx context.Context, username string) (string, int, error)
	Role(ctx context.Context, userId int) (string, error)
}

func NewService(repo Repository) Service {
	return &serviceImplementation{repo: repo}
}

type serviceImplementation struct {
	repo Repository
}

func (s *serviceImplementation) GetRole(ctx context.Context, userId int) (string, error) {
	return s.repo.Role(ctx, userId)
}

func (s *serviceImplementation) CreateUser(ctx context.Context, firstName, lastName, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	err = s.repo.Create(ctx, firstName, lastName, email, string(hashedPassword), "USER", 0, nil)
	return err
}

func (s *serviceImplementation) UpdateDeleteDate(ctx context.Context, userId int) error {
	return s.repo.UpdateDeleteDate(ctx, userId, time.Now().AddDate(0, 1, 0))
}

func (s *serviceImplementation) GetUserPasswordAndId(ctx context.Context, username string) (string, int, error) {
	return s.repo.PasswordAndId(ctx, username)
}
