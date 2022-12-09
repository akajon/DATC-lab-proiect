package users

type Repository interface {
}

func NewService(repo Repository) Service {
	return &serviceImplementation{repo: repo}
}

type serviceImplementation struct {
	repo Repository
}
