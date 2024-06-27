package services

import (
	"context"
	"github.com/matheusvidal21/product-recommendation-service/application/repositories"
	"github.com/matheusvidal21/product-recommendation-service/domain/models"
)

type CategoryServiceInterface interface {
	GetAllCategories() ([]models.CategoryDomain, error)
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

func (s *CategoryService) GetAllCategories() ([]models.CategoryDomain, error) {
	return s.repo.FindAll()
}

func (s *CategoryService) GetCategoryByID(id string) (*models.CategoryDomain, error) {
	return s.repo.FindByID(id)
}

func (s *CategoryService) CreateCategory(name, description string) (*models.CategoryDomain, error) {
	category := models.NewCategoryDomain("", name, description)
	return s.repo.Create(category)
}

func (s *CategoryService) UpdateCategory(id, name, description string) (*models.CategoryDomain, error) {
	category := models.NewCategoryDomain(id, name, description)
	return s.repo.Update(category)
}

func (s *CategoryService) DeleteCategory(id string) error {
	return s.repo.Delete(id)
}
