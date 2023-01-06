package alerts

import (
	"context"
	"database/sql"
	"strconv"
	"time"
)

func NewRepository(db *sql.DB) Repository {
	return &repositoryImpl{db: db}
}

type repositoryImpl struct {
	db *sql.DB
}

func (r *repositoryImpl) Delete(ctx context.Context, alertId int) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM dbo.alerts WHERE alert_id = @alert_id`, sql.Named("alert_id", alertId))

	return err
}

func (r *repositoryImpl) RewardUser(ctx context.Context, userId, taxReduction int) error {
	_, err := r.db.ExecContext(ctx, `UPDATE dbo.users SET tax_reduction = tax_reduction + @tax_reduction WHERE user_id = @user_id`, sql.Named("tax_reduction", taxReduction), sql.Named("user_id", userId))

	return err
}

func (r *repositoryImpl) Get(ctx context.Context, alertId int) (*AlertReward, error) {
	var alert AlertReward

	err := r.db.QueryRowContext(ctx, `SELECT danger_id, alert_owner_id, users FROM dbo.alerts WHERE alert_id = @alert_id`, sql.Named("alert_id", alertId)).Scan(&alert.DangerId, &alert.OwnerId, &alert.Users)

	if err != nil {
		return nil, err
	}

	return &alert, nil
}

func (r *repositoryImpl) GetDangerGrade(ctx context.Context, dangerId int) (int, error) {
	var dangerGrade int

	err := r.db.QueryRowContext(ctx, `SELECT grade FROM dbo.dangers WHERE danger_id = @danger_id`, sql.Named("danger_id", dangerId)).Scan(&dangerGrade)

	if err != nil {
		return 0, err
	}

	return dangerGrade, nil
}

func (r *repositoryImpl) Verify(ctx context.Context, dangerId int, latitude, longitude float32) (int, error) {
	var alertId int
	err := r.db.QueryRowContext(ctx, `SELECT alert_id FROM dbo.alerts WHERE danger_id = @danger_id AND latitude = @latitude AND longitude = @longitude`,
		sql.Named("danger_id", dangerId), sql.Named("latitude", latitude), sql.Named("longitude", longitude)).Scan(&alertId)

	if err != nil {
		return 0, err
	}

	return alertId, nil
}

func (r *repositoryImpl) Create(ctx context.Context, userId, dangerId int, latitude, longitude float32) error {
	_, err := r.db.ExecContext(ctx, `INSERT INTO dbo.alerts (alert_owner_id, danger_id, users, latitude, longitude, date) VALUES (@alert_owner_id, @danger_id, @users, @latitude, @longitude, @date)`,
		sql.Named("alert_owner_id", userId), sql.Named("danger_id", dangerId), sql.Named("users", ""), sql.Named("latitude", latitude), sql.Named("longitude", longitude), sql.Named("date", time.Now()))

	return err
}

func (r *repositoryImpl) AddUser(ctx context.Context, userId, alertId int) error {
	_, err := r.db.ExecContext(ctx, `UPDATE dbo.alerts SET users = CONCAT(users,@user) WHERE alert_id = @alert_id`,
		sql.Named("alert_id", alertId), sql.Named("user", strconv.Itoa(userId)+" "))

	return err
}
