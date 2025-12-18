package service

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/example/user-age-api/internal/models"
	"github.com/example/user-age-api/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, req models.CreateUserRequest) (models.User, error)
	GetUser(ctx context.Context, id int64) (models.User, error)
	ListUsers(ctx context.Context) ([]models.User, error)
	UpdateUser(ctx context.Context, id int64, req models.UpdateUserRequest) (models.User, error)
	DeleteUser(ctx context.Context, id int64) error
}

type userService struct {
	repo      repository.UserRepository
	validator *validator.Validate
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo:      repo,
		validator: validator.New(),
	}
}

func (s *userService) CreateUser(ctx context.Context, req models.CreateUserRequest) (models.User, error) {
	if err := s.validator.Struct(req); err != nil {
		return models.User{}, err
	}

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return models.User{}, err
	}

	u, err := s.repo.CreateUser(ctx, req.Name, dob)
	if err != nil {
		return models.User{}, err
	}

	return models.User{ID: u.ID, Name: u.Name, DOB: u.DOB}, nil
}

func (s *userService) GetUser(ctx context.Context, id int64) (models.User, error) {
	u, err := s.repo.GetUser(ctx, id)
	if err != nil {
		return models.User{}, err
	}
	return models.User{ID: u.ID, Name: u.Name, DOB: u.DOB}, nil
}

func (s *userService) ListUsers(ctx context.Context) ([]models.User, error) {
	dbUsers, err := s.repo.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	users := make([]models.User, 0, len(dbUsers))
	for _, u := range dbUsers {
		users = append(users, models.User{ID: u.ID, Name: u.Name, DOB: u.DOB})
	}
	return users, nil
}

func (s *userService) UpdateUser(ctx context.Context, id int64, req models.UpdateUserRequest) (models.User, error) {
	if err := s.validator.Struct(req); err != nil {
		return models.User{}, err
	}

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return models.User{}, err
	}

	u, err := s.repo.UpdateUser(ctx, id, req.Name, dob)
	if err != nil {
		return models.User{}, err
	}

	return models.User{ID: u.ID, Name: u.Name, DOB: u.DOB}, nil
}

func (s *userService) DeleteUser(ctx context.Context, id int64) error {
	return s.repo.DeleteUser(ctx, id)
}


