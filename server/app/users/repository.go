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

func (r repositoryImpl) Create(ctx context.Context, firstName, lastName, email, hashedPassword, role string, taxReduction int, deletionDate *time.Time) (*CreateUserResponse, error) {
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

	if err != nil {
		return nil, err
	}

	newUser := CreateUserResponse{
		Id:           id,
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		Password:     hashedPassword,
		TaxReduction: taxReduction,
		Role:         role,
		DeletionDate: deletionDate,
	}

	return &newUser, nil
}

func (r repositoryImpl) UpdateDeleteDate(ctx context.Context, userId int, deleteDate time.Time) (*UpdateDeleteDateResponse, error) {
	_, err := r.db.ExecContext(ctx, `UPDATE dbo.users SET deletion_date = @deletion_date where user_id = @user_id`,
		sql.Named("deletion_date", deleteDate),
		sql.Named("user_id", userId))

	if err != nil {
		return nil, err
	}
	return &UpdateDeleteDateResponse{Id: userId, DeletionDate: deleteDate}, nil
}
