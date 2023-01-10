package users

import (
	"context"
	"database/sql"
	"time"
)

func NewRepository(db *sql.DB) Repository {
	return &repositoryImpl{db: db}
}

type repositoryImpl struct {
	db *sql.DB
}

func (r *repositoryImpl) Create(ctx context.Context, firstName, lastName, email, hashedPassword, role string, taxReduction int, deletionDate *time.Time) error {
	var id int

	err := r.db.QueryRowContext(ctx, `INSERT INTO dbo.users (first_name, last_name, email, passw, tax_reduction, rol, deletion_date) OUTPUT inserted.user_id
						VALUES (@first_name, @last_name, @email, @passw, @tax_reduction, @rol, @deletion_date)`,
		sql.Named("first_name", firstName),
		sql.Named("last_name", lastName),
		sql.Named("email", email),
		sql.Named("passw", hashedPassword),
		sql.Named("tax_reduction", taxReduction),
		sql.Named("rol", role),
		sql.Named("deletion_date", deletionDate)).Scan(&id)

	return err
}

func (r *repositoryImpl) UpdateDeleteDate(ctx context.Context, userId int, deleteDate time.Time) error {
	_, err := r.db.ExecContext(ctx, `UPDATE dbo.users SET deletion_date = @deletion_date where user_id = @user_id`,
		sql.Named("deletion_date", deleteDate),
		sql.Named("user_id", userId))

	return err
}

func (r *repositoryImpl) PasswordAndId(ctx context.Context, username string) (string, int, error) {
	var password string
	var id int
	err := r.db.QueryRowContext(ctx, `SELECT passw, user_id FROM dbo.users WHERE email = @email`, sql.Named("email", username)).Scan(&password, &id)

	if err != nil {
		return "", 0, err
	}
	return password, id, nil
}

func (r *repositoryImpl) Role(ctx context.Context, userId int) (string, error) {
	var role string

	err := r.db.QueryRowContext(ctx, `SELECT rol FROM dbo.users WHERE user_id = @user_id`, sql.Named("user_id", userId)).Scan(&role)
	if err != nil {
		return "", err
	}

	return role, nil
}

func (r *repositoryImpl) Get(ctx context.Context, userId int) (*SignInResponse, error) {
	var user SignInResponse

	user.Id = userId
	err := r.db.QueryRowContext(ctx, `SELECT first_name, last_name, email, tax_reduction, rol FROM dbo.users WHERE user_id = @user_id`, sql.Named("user_id", userId)).
		Scan(&user.FirstName, &user.LastName, &user.Email, &user.TaxReduction, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
