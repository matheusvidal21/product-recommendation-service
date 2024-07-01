package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/matheusvidal21/product-recommendation-service/application/repositories"
	"github.com/matheusvidal21/product-recommendation-service/domain/models"
	"github.com/matheusvidal21/product-recommendation-service/domain/models/dtos"
	logger "github.com/matheusvidal21/product-recommendation-service/framework/config/logging"
)

type CategoryServiceInterface interface {
	GetAllCategories() ([]dtos.CategoryDTO, error)
	GetCategoryByID(id string) (*models.CategoryDomain, error)
	CreateCategory(name, description string) (*models.CategoryDomain, error)
	UpdateCategory(id, name, description string) (*models.CategoryDomain, error)
	DeleteCategory(id string) error
}

type CategoryService struct {
	repo repositories.CategoryRepositoryInterface
	ctx  context.Context
}

func NewCategoryService(repo repositories.CategoryRepositoryInterface, ctx context.Context) CategoryServiceInterface {
	return &CategoryService{
		repo: repo,
		ctx:  ctx,
	}
}

func (s *CategoryService) GetAllCategories() ([]dtos.CategoryDTO, error) {
	logger.Info("Getting all categories services")
	cat := []dtos.CategoryDTO{}

	categories, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	for _, c := range categories {
		cat = append(cat, dtos.CategoryToDTO(c))
	}
	return cat, nil
}

func (s *CategoryService) GetCategoryByID(id string) (*models.CategoryDomain, error) {
	logger.Info("Getting category by id services")
	return s.repo.FindByID(id)
}

func (s *CategoryService) CreateCategory(name, description string) (*models.CategoryDomain, error) {
	logger.Info("Creating category services")
	id := uuid.New().String()
	category := models.NewCategoryDomain(id, name, description)
	return s.repo.Create(category)
}

func (s *CategoryService) UpdateCategory(id, name, description string) (*models.CategoryDomain, error) {
	logger.Info("Updating category services")
	category := models.NewCategoryDomain(id, name, description)
	return s.repo.Update(id, category)
}

func (s *CategoryService) DeleteCategory(id string) error {
	logger.Info("Deleting category services")
	return s.repo.Delete(id)
}
