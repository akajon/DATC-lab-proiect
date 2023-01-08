package alerts

import (
	"context"
	"errors"
	"strconv"
	"strings"
)

type Repository interface {
	Verify(ctx context.Context, dangerId int, latitude, longitude float32) (int, error)
	Create(ctx context.Context, userId, dangerId int, latitude, longitude float32) error
	AddUser(ctx context.Context, userId, alertId int) error
	Delete(ctx context.Context, alertId int) error
	RewardUser(ctx context.Context, userId, taxReduction int) error
	Get(ctx context.Context, alertId int) (*AlertReward, error)
	GetDangerGrade(ctx context.Context, dangerId int) (int, error)
	GetAll(ctx context.Context) ([]AlertGetResponse, error)
}

func NewService(repo Repository) Service {
	return &serviceImplementation{repo: repo}
}

type serviceImplementation struct {
	repo Repository
}

func (s *serviceImplementation) DeleteAlert(ctx context.Context, alertId int) error {
	alert, err := s.repo.Get(ctx, alertId)
	if err != nil {
		return err
	}

	dangerGrade, err := s.repo.GetDangerGrade(ctx, alert.DangerId)
	if err != nil {
		return err
	}

	var users []int
	usersStrId := strings.Split(alert.Users, " ")
	usersStrId = usersStrId[:len(usersStrId)-1]
	for _, strId := range usersStrId {
		userId, err := strconv.Atoi(strId)
		if err != nil {
			return err
		}
		users = append(users, userId)
	}

	users = append(users, alert.OwnerId)
	for _, u := range users {
		err = s.repo.RewardUser(ctx, u, dangerGrade)
		if err != nil {
			return err
		}
	}

	err = s.repo.Delete(ctx, alertId)
	if err != nil {
		return err
	}

	return nil
}

func (s *serviceImplementation) VerifyAlert(ctx context.Context, dangerId int, latitude, longitude float32) (int, error) {
	return s.repo.Verify(ctx, dangerId, latitude, longitude)
}

func (s *serviceImplementation) CreateAlert(ctx context.Context, userId, dangerId int, latitude, longitude float32) error {
	return s.repo.Create(ctx, userId, dangerId, latitude, longitude)
}

func (s *serviceImplementation) AddUserToAlert(ctx context.Context, userId, alertId int) error {
	alert, err := s.repo.Get(ctx, alertId)
	if err != nil {
		return err
	}

	var users []int
	usersStrId := strings.Split(alert.Users, " ")
	usersStrId = usersStrId[:len(usersStrId)-1]
	for _, strId := range usersStrId {
		userId, err := strconv.Atoi(strId)
		if err != nil {
			return err
		}
		users = append(users, userId)
	}

	users = append(users, alert.OwnerId)
	for _, u := range users {
		if u == userId {
			return errors.New("user already reported this alert")
		}
	}

	return s.repo.AddUser(ctx, userId, alertId)
}

func (s *serviceImplementation) GetAlerts(ctx context.Context) ([]AlertGetResponse, error) {
	return s.repo.GetAll(ctx)
}
