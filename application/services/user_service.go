package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/matheusvidal21/product-recommendation-service/application/repositories"
	"github.com/matheusvidal21/product-recommendation-service/domain/models"
	"github.com/matheusvidal21/product-recommendation-service/domain/models/dtos"
	logger "github.com/matheusvidal21/product-recommendation-service/framework/config/logging"
)

type UserServiceInterface interface {
	GetAllUsers() ([]dtos.UserResponseDTO, error)
	GetUserByID(id string) (*models.UserDomain, error)
	CreateUser(name, email, password string) (*models.UserDomain, error)
	UpdateUser(id, name, email, password string) (*models.UserDomain, error)
	DeleteUser(id string) error
}

type UserService struct {
	repo repositories.UserRepositoryInterface
	ctx  context.Context
}

func NewUserService(repo repositories.UserRepositoryInterface, ctx context.Context) UserServiceInterface {
	return &UserService{
		repo: repo,
		ctx:  ctx,
	}
}

func (s *UserService) GetAllUsers() ([]dtos.UserResponseDTO, error) {
	logger.Info("Getting all users")
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var usersDTO []dtos.UserResponseDTO
	for _, user := range users {
		usersDTO = append(usersDTO, dtos.UserResponseDTO{
			ID:    user.GetID(),
			Name:  user.GetName(),
			Email: user.GetEmail(),
		})
	}

	return usersDTO, nil
}

func (s *UserService) GetUserByID(id string) (*models.UserDomain, error) {
	logger.Info("Getting user by id services")
	return s.repo.FindByID(id)
}

func (s *UserService) CreateUser(name, email, password string) (*models.UserDomain, error) {
	logger.Info("Creating user services")
	id := uuid.New().String()
	user := models.NewUserDomain(id, name, email, password)
	return s.repo.Create(user)
}

func (s *UserService) UpdateUser(id, name, email, password string) (*models.UserDomain, error) {
	logger.Info("Updating user services")
	user := models.NewUserDomain(id, name, email, password)
	return s.repo.Update(id, user)
}

func (s *UserService) DeleteUser(id string) error {
	logger.Info("Deleting user services")
	return s.repo.Delete(id)
}
