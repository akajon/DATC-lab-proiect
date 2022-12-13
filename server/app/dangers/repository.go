package dangers

import (
	"context"
	"database/sql"
)

func NewRepository(db *sql.DB) Repository {
	return &repositoryImpl{db: db}
}

type repositoryImpl struct {
	db *sql.DB
}

func (r repositoryImpl) Create(ctx context.Context, category, name, description string, grade int) (*CreateDangerResponse, error) {
	err := r.db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	var id int
	err = r.db.QueryRow(`INSERT INTO dbo.dangers (category, name, description, grade) OUTPUT inserted.danger_id
							   VALUES (@category, @name, @description, @grade)`,
		sql.Named("category", category),
		sql.Named("name", name),
		sql.Named("description", description),
		sql.Named("grade", grade)).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &CreateDangerResponse{Id: id, Category: category, Name: name, Description: description, Grade: grade}, nil
}
