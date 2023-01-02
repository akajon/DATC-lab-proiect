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

func (r repositoryImpl) Create(ctx context.Context, category, name, description string, grade int) error {
	var id int

	err := r.db.QueryRowContext(ctx, `INSERT INTO dbo.dangers (category, name, description, grade) OUTPUT inserted.danger_id
							   VALUES (@category, @name, @description, @grade)`,
		sql.Named("category", category),
		sql.Named("name", name),
		sql.Named("description", description),
		sql.Named("grade", grade)).Scan(&id)

	return err
}

func (r repositoryImpl) Delete(ctx context.Context, dangerId int) error {

	_, err := r.db.ExecContext(ctx, "DELETE FROM dbo.dangers WHERE danger_id = @danger_id", sql.Named("danger_id", dangerId))

	return err
}
