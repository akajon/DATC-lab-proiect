package alerts

import "context"

type Repository interface {
	Verify(ctx context.Context, dangerId int, latitude, longitude float32) (error, int)
	Create(ctx context.Context, userId, dangerId int, latitude, longitude float32) error
	AddUser(ctx context.Context, userId, alertId int) error
}

func NewService(repo Repository) Service {
	return &serviceImplementation{repo: repo}
}

type serviceImplementation struct {
	repo Repository
}

func (s serviceImplementation) VerifyAlert(ctx context.Context, dangerId int, latitude, longitude float32) (error, int) {
	return s.repo.Verify(ctx, dangerId, latitude, longitude)
}

func (s serviceImplementation) CreateAlert(ctx context.Context, userId, dangerId int, latitude, longitude float32) error {
	return s.repo.Create(ctx, userId, dangerId, latitude, longitude)
}

func (s serviceImplementation) AddUserToAlert(ctx context.Context, userId, alertId int) error {
	return s.repo.AddUser(ctx, userId, alertId)
}
