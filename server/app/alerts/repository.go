package alerts

import "database/sql"

func NewRepository(db *sql.DB) Repository {
	return &repositoryImpl{db: db}
}

type repositoryImpl struct {
	db *sql.DB
}
