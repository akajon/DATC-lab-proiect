package dangers

import "context"

type Repository interface {
	Create(ctx context.Context, category, name, description string, grade int) (*CreateDangerResponse, error)
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
