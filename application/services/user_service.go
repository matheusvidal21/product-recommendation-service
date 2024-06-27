package services

import (
	"context"
	"github.com/matheusvidal21/product-recommendation-service/application/repositories"
	"github.com/matheusvidal21/product-recommendation-service/domain/models"
)

type UserServiceInterface interface {
	GetAllUsers() ([]models.UserDomain, error)
	GetUserByID(id string) (*models.UserDomain, error)
	CreateUser(name, email, password string) (*models.UserDomain, error)
	UpdateUser(name, email, password string) (*models.UserDomain, error)
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

func (s *UserService) GetAllUsers() ([]models.UserDomain, error) {
	return s.repo.FindAll()
}

func (s *UserService) GetUserByID(id string) (*models.UserDomain, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) CreateUser(name, email, password string) (*models.UserDomain, error) {
	user := models.NewUser("", name, email, password)
	return s.repo.Create(user)
}

func (s *UserService) UpdateUser(name, email, password string) (*models.UserDomain, error) {
	user := models.NewUser("", name, email, password)
	return s.repo.Update(user)
}

func (s *UserService) DeleteUser(id string) error {
	return s.repo.Delete(id)
}
