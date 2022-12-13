package dangers

import "context"

type Repository interface {
	Create(ctx context.Context, category, name, description string, grade int) (*CreateDangerResponse, error)
	Delete(ctx context.Context, dangerId int) error
}

func NewService(repo Repository) Service {
	return &serviceImplementation{repo: repo}
}

type serviceImplementation struct {
	repo Repository
}

func (s serviceImplementation) CreateDanger(ctx context.Context, category, name, description string, grade int) (*CreateDangerResponse, error) {
	return s.repo.Create(ctx, category, name, description, grade)
}

func (s serviceImplementation) DeleteDanger(ctx context.Context, dangerId int) error {
	return s.repo.Delete(ctx, dangerId)
}
