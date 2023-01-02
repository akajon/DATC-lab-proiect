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

func (r repositoryImpl) Verify(ctx context.Context, dangerId int, latitude, longitude float32) (error, int) {
	var alertId int
	err := r.db.QueryRowContext(ctx, `SELECT alert_id FROM dbo.alerts WHERE danger_id = @danger_id AND latitude = @latitude AND longitude = @longitude`,
		sql.Named("danger_id", dangerId), sql.Named("latitude", latitude), sql.Named("longitude", longitude)).Scan(&alertId)

	if err != nil {
		return err, 0
	}

	return nil, alertId
}

func (r repositoryImpl) Create(ctx context.Context, userId, dangerId int, latitude, longitude float32) error {
	_, err := r.db.ExecContext(ctx, `INSERT INTO dbo.alerts (alert_owner_id, danger_id, users, latitude, longitude, date) VALUES (@alert_owner_id, @danger_id, @users, @latitude, @longitude, @date)`,
		sql.Named("alert_owner_id", userId), sql.Named("danger_id", dangerId), sql.Named("users", ""), sql.Named("latitude", latitude), sql.Named("longitude", longitude), sql.Named("date", time.Now()))

	return err
}

func (r repositoryImpl) AddUser(ctx context.Context, userId, alertId int) error {
	_, err := r.db.ExecContext(ctx, `UPDATE dbo.alerts SET users = CONCAT(users,@user) WHERE alert_id = @alert_id`,
		sql.Named("alert_id", alertId), sql.Named("user", strconv.Itoa(userId)+" "))

	return err
}
