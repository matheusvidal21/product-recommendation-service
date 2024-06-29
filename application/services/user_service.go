package services

import (
	"context"
	"github.com/matheusvidal21/product-recommendation-service/application/repositories"
	"github.com/matheusvidal21/product-recommendation-service/domain/models"
	"github.com/matheusvidal21/product-recommendation-service/domain/models/dtos"
)

type UserServiceInterface interface {
	GetAllUsers() ([]dtos.UserGetAllDTO, error)
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

func (s *UserService) GetAllUsers() ([]dtos.UserGetAllDTO, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var usersDTO []dtos.UserGetAllDTO
	for _, user := range users {
		usersDTO = append(usersDTO, dtos.UserGetAllDTO{
			ID:    user.GetID(),
			Name:  user.GetName(),
			Email: user.GetEmail(),
		})
	}

	return usersDTO, nil
}

func (s *UserService) GetUserByID(id string) (*models.UserDomain, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) CreateUser(name, email, password string) (*models.UserDomain, error) {
	user := models.NewUserDomain(name, email, password)
	return s.repo.Create(user)
}

func (s *UserService) UpdateUser(id, name, email, password string) (*models.UserDomain, error) {
	user := models.NewUserWithId(id, name, email, password)
	return s.repo.Update(user)
}

func (s *UserService) DeleteUser(id string) error {
	return s.repo.Delete(id)
}
